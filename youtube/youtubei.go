package youtube

import (
   "bytes"
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "mime"
   "net/http"
   "net/url"
   "path"
   "strings"
   "time"
)

func (c Context) PlayerHeader(head http.Header, id string) (*Player, error) {
   var body struct {
      Context Context `json:"context"`
      RacyCheckOK bool `json:"racyCheckOk,omitempty"`
      VideoID string `json:"videoId"`
   }
   body.Context = c
   body.VideoID = id
   if head.Get("Authorization") != "" {
      body.RacyCheckOK = true // Cr381pDsSsA
   }
   buf, err := encode(body)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", origin + "/youtubei/v1/player", buf)
   if err != nil {
      return nil, err
   }
   req.Header = head
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   play := new(Player)
   if err := json.NewDecoder(res.Body).Decode(play); err != nil {
      return nil, err
   }
   return play, nil
}

const origin = "https://www.youtube.com"

var googAPI = http.Header{
   "X-Goog-Api-Key": {"AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"},
}

// https://youtube.com/shorts/9Vsdft81Q6w
// https://youtube.com/watch?v=XY-hOqcPGCY
func VideoID(address string) (string, error) {
   parse, err := url.Parse(address)
   if err != nil {
      return "", err
   }
   v := parse.Query().Get("v")
   if v != "" {
      return v, nil
   }
   return path.Base(parse.Path), nil
}

func encode(val any) (*bytes.Buffer, error) {
   buf := new(bytes.Buffer)
   enc := json.NewEncoder(buf)
   enc.SetIndent("", " ")
   err := enc.Encode(val)
   if err != nil {
      return nil, err
   }
   return buf, nil
}

type Client struct {
   Name string `json:"clientName"`
   Screen string `json:"clientScreen,omitempty"`
   Version string `json:"clientVersion"`
}

type Context struct {
   Client Client `json:"client"`
   ThirdParty *ThirdParty `json:"thirdParty,omitempty"`
}

var Android = Context{
   Client: Client{Name: "ANDROID", Version: "17.09.33"},
}

// HsUATh_Nc2U
var Embed = Context{
   Client: Client{Name: "ANDROID", Screen: "EMBED", Version: "17.09.33"},
   ThirdParty: &ThirdParty{EmbedURL: origin},
}

var Mweb = Context{
   Client: Client{Name: "MWEB", Version: "2.20211109.01.00"},
}

func (c Context) Player(id string) (*Player, error) {
   return c.PlayerHeader(googAPI, id)
}

func (c Context) Search(query string) (*Search, error) {
   var body struct {
      Context Context `json:"context"`
      Params string `json:"params"`
      Query string `json:"query"`
   }
   body.Context = c
   filter := NewFilter()
   filter.Type(Type["Video"])
   param := NewParams()
   param.Filter(filter)
   body.Params = param.Encode()
   body.Query = query
   buf := new(bytes.Buffer)
   if err := json.NewEncoder(buf).Encode(body); err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", origin + "/youtubei/v1/search", buf)
   if err != nil {
      return nil, err
   }
   req.Header = googAPI
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   search := new(Search)
   if err := json.NewDecoder(res.Body).Decode(search); err != nil {
      return nil, err
   }
   return search, nil
}

type Item struct {
   CompactVideoRenderer *struct {
      Title struct {
         Runs []struct {
            Text string
         }
      }
      VideoID string
   }
}

type Player struct {
   PlayabilityStatus struct {
      Status string // "OK", "LOGIN_REQUIRED"
      Reason string // "", "Sign in to confirm your age"
   }
   VideoDetails struct {
      VideoID string
      LengthSeconds int64 `json:"lengthSeconds,string"`
      ViewCount int64 `json:"viewCount,string"`
      Author string
      Title string
      ShortDescription string
   }
   Microformat struct {
      PlayerMicroformatRenderer struct {
         PublishDate string // 2013-06-11
      }
   }
   StreamingData StreamingData
}

func (p Player) Base() string {
   return mech.Clean(p.VideoDetails.Author + "-" + p.VideoDetails.Title)
}

func (p Player) Date() (time.Time, error) {
   value := p.Microformat.PlayerMicroformatRenderer.PublishDate
   return time.Parse("2006-01-02", value)
}

func (p Player) Format(f fmt.State, verb rune) {
   fmt.Fprintln(f, p.Status())
   fmt.Fprintln(f, "VideoID:", p.VideoDetails.VideoID)
   fmt.Fprintln(f, "Length:", p.VideoDetails.LengthSeconds)
   fmt.Fprintln(f, "ViewCount:", p.VideoDetails.ViewCount)
   fmt.Fprintln(f, "Author:", p.VideoDetails.Author)
   fmt.Fprintln(f, "Title:", p.VideoDetails.Title)
   date := p.Microformat.PlayerMicroformatRenderer.PublishDate
   if date != "" {
      fmt.Fprintln(f, "Date:", date)
   }
   for _, form := range p.StreamingData.AdaptiveFormats {
      fmt.Fprintln(f)
      form.Format(f, verb)
   }
}

func (p Player) Status() string {
   var buf strings.Builder
   buf.WriteString("Status: ")
   buf.WriteString(p.PlayabilityStatus.Status)
   if p.PlayabilityStatus.Reason != "" {
      buf.WriteString("\nReason: ")
      buf.WriteString(p.PlayabilityStatus.Reason)
   }
   return buf.String()
}

type Search struct {
   Contents struct {
      SectionListRenderer struct {
         Contents []struct {
            ItemSectionRenderer *struct {
               Contents []Item
            }
         }
      }
   }
}

func (s Search) Items() []Item {
   var items []Item
   for _, sect := range s.Contents.SectionListRenderer.Contents {
      if sect.ItemSectionRenderer != nil {
         for _, item := range sect.ItemSectionRenderer.Contents {
            if item.CompactVideoRenderer != nil {
               items = append(items, item)
            }
         }
      }
   }
   return items
}

type StreamingData struct {
   AdaptiveFormats []Format
}

func (s StreamingData) Len() int {
   return len(s.AdaptiveFormats)
}

func (s *StreamingData) MediaType() error {
   for i, form := range s.AdaptiveFormats {
      typ, param, err := mime.ParseMediaType(form.MimeType)
      if err != nil {
         return err
      }
      param["codecs"], _, _ = strings.Cut(param["codecs"], ".")
      s.AdaptiveFormats[i].MimeType = mime.FormatMediaType(typ, param)
   }
   return nil
}

func (s StreamingData) Swap(i, j int) {
   swap := s.AdaptiveFormats[i]
   s.AdaptiveFormats[i] = s.AdaptiveFormats[j]
   s.AdaptiveFormats[j] = swap
}

type ThirdParty struct {
   EmbedURL string `json:"embedUrl"`
}
