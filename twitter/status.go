package twitter

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "path"
   "strconv"
)

const bearer =
   "AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs=" +
   "1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA"

var LogLevel format.LogLevel

type Guest struct {
   Guest_Token string
}

func NewGuest() (*Guest, error) {
   req, err := http.NewRequest(
      "POST", "https://api.twitter.com/1.1/guest/activate.json", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Authorization", "Bearer " + bearer)
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   guest := new(Guest)
   if err := json.NewDecoder(res.Body).Decode(guest); err != nil {
      return nil, err
   }
   return guest, nil
}

func (g Guest) Status(id int64) (*Status, error) {
   buf := []byte("https://api.twitter.com/1.1/statuses/show/")
   buf = strconv.AppendInt(buf, id, 10)
   buf = append(buf, ".json?tweet_mode=extended"...)
   req, err := http.NewRequest("GET", string(buf), nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + bearer},
      "X-Guest-Token": {g.Guest_Token},
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   stat := new(Status)
   if err := json.NewDecoder(res.Body).Decode(stat); err != nil {
      return nil, err
   }
   return stat, nil
}

type Media struct {
   Media_URL string
   Original_Info struct {
      Width int64
      Height int64
   }
   Video_Info struct {
      Variants []Variant
   }
}

func (m Media) String() string {
   buf := []byte("Width: ")
   buf = strconv.AppendInt(buf, m.Original_Info.Width, 10)
   buf = append(buf, "\nHeight: "...)
   buf = strconv.AppendInt(buf, m.Original_Info.Height, 10)
   for _, vari := range m.Variants() {
      buf = append(buf, '\n')
      buf = append(buf, vari.String()...)
   }
   return string(buf)
}

func (m Media) Variants() []Variant {
   var varis []Variant
   for _, vari := range m.Video_Info.Variants {
      if vari.Content_Type != "application/x-mpegURL"{
         varis = append(varis, vari)
      }
   }
   return varis
}

type Status struct {
   Created_At string
   User struct {
      Screen_Name string
      Name string
   }
   Full_Text string
   Extended_Entities struct {
      Media []Media
   }
}

func (s Status) Base(id int64) string {
   var buf []byte
   buf = append(buf, s.User.Screen_Name...)
   buf = append(buf, '-')
   buf = strconv.AppendInt(buf, id, 10)
   return string(buf)
}

func (s Status) String() string {
   var buf []byte
   buf = append(buf, "Screen Name: "...)
   buf = append(buf, s.User.Screen_Name...)
   buf = append(buf, "\nName: "...)
   buf = append(buf, s.User.Name...)
   buf = append(buf, "\nCreated: "...)
   buf = append(buf, s.Created_At...)
   buf = append(buf, "\nText: "...)
   buf = append(buf, s.Full_Text...)
   for _, media := range s.Extended_Entities.Media {
      buf = append(buf, '\n')
      buf = append(buf, media.String()...)
   }
   return string(buf)
}

type Variant struct {
   Bitrate int64
   Content_Type string
   URL string
}

func (v Variant) Ext() (string, error) {
   addr, err := url.Parse(v.URL)
   if err != nil {
      return "", err
   }
   return path.Ext(addr.Path), nil
}

func (v Variant) String() string {
   var buf []byte
   buf = append(buf, "Bitrate:"...)
   buf = strconv.AppendInt(buf, v.Bitrate, 10)
   buf = append(buf, " URL:"...)
   buf = append(buf, v.URL...)
   return string(buf)
}
