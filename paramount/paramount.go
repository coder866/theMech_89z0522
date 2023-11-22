package paramount

import (
   "encoding/xml"
   "github.com/89z/format"
   "github.com/89z/mech"
   "net/http"
   "net/url"
   "strconv"
)

const (
   aid = 2198311517
   sid = "dJ5BDC"
)

var LogLevel format.LogLevel

type Media struct {
   Body struct {
      Seq  struct {
         Video struct {
            Title string `xml:"title,attr"`
            Src string `xml:"src,attr"`
         } `xml:"video"`
      } `xml:"seq"`
   } `xml:"body"`
}

func NewMedia(guid string) (*Media, error) {
   buf := []byte("https://link.theplatform.com/s/")
   buf = append(buf, sid...)
   buf = append(buf, "/media/guid/"...)
   buf = strconv.AppendInt(buf, aid, 10)
   buf = append(buf, '/')
   buf = append(buf, guid...)
   req, err := http.NewRequest("GET", string(buf), nil)
   if err != nil {
      return nil, err
   }
   // We need "MPEG4", otherwise you get a "EXT-X-KEY" with "skd" scheme:
   req.URL.RawQuery = url.Values{
      "format": {"SMIL"},
      "formats": {"MPEG4,M3U"},
   }.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   med := new(Media)
   if err := xml.NewDecoder(res.Body).Decode(med); err != nil {
      return nil, err
   }
   return med, nil
}

func (m Media) Base() string {
   return mech.Clean(m.Body.Seq.Video.Title)
}
