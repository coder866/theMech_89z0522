package tiktok

import (
   "encoding/json"
   "github.com/89z/format"
   "github.com/89z/mech"
   "net/http"
   "strconv"
   "time"
)

const origin = "http://api2.musical.ly"

var LogLevel format.LogLevel

type Detail struct {
   Author struct {
      Unique_ID string
   }
   Aweme_ID string
   Create_Time int64
   Video struct {
      Duration int64
      Play_Addr struct {
         Width int64
         Height int64 // this is invalid
         URL_List []string
      }
   }
}

func (d Detail) Base() string {
   return mech.Clean(d.Author.Unique_ID + "-" + d.Aweme_ID)
}

func NewDetail(id int64) (*Detail, error) {
   req, err := http.NewRequest("GET", origin + "/aweme/v1/aweme/detail/", nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "aweme_id=" + strconv.FormatInt(id, 10)
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var aweme struct {
      Aweme_Detail Detail
   }
   if err := json.NewDecoder(res.Body).Decode(&aweme); err != nil {
      return nil, err
   }
   return &aweme.Aweme_Detail, nil
}

func (d Detail) Duration() time.Duration {
   return time.Duration(d.Video.Duration) * time.Millisecond
}

func (d Detail) String() string {
   var buf []byte
   buf = append(buf, "ID: "...)
   buf = append(buf, d.Aweme_ID...)
   buf = append(buf, "\nAuthor: "...)
   buf = append(buf, d.Author.Unique_ID...)
   buf = append(buf, "\nCreate_Time: "...)
   buf = append(buf, d.Time().String()...)
   buf = append(buf, "\nDuration: "...)
   buf = append(buf, d.Duration().String()...)
   buf = append(buf, "\nWidth: "...)
   buf = strconv.AppendInt(buf, d.Video.Play_Addr.Width, 10)
   buf = append(buf, "\nHeight: "...)
   buf = strconv.AppendInt(buf, d.Video.Play_Addr.Height, 10)
   for _, addr := range d.Video.Play_Addr.URL_List {
      buf = append(buf, "\nURL: "...)
      buf = append(buf, addr...)
   }
   return string(buf)
}

func (d Detail) Time() time.Time {
   return time.Unix(d.Create_Time, 0)
}

func (d Detail) URL() string {
   for _, addr := range d.Video.Play_Addr.URL_List {
      return addr
   }
   return ""
}
