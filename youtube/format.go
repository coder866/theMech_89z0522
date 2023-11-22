package youtube

import (
   "fmt"
   "github.com/89z/format"
   "io"
   "mime"
   "net/http"
   "time"
)

const partLength = 10_000_000

var LogLevel format.LogLevel

// averageBitrate is a better marker for quality than bitrate. For example, if
// you look a video:
// 7WTEB7Qbt4U
// you get this:
//
// itag | bitrate | averageBitrate | contentLength
// -----|---------|----------------|--------------
// 136  | 1038025 | 286687         | 6891870
// 247  | 1192816 | 513601         | 12346788
// 398  | 1117347 | 349310         | 8397292
type Format struct {
   AudioQuality string
   AverageBitrate int
   ContentLength int64 `json:"contentLength,string"`
   Height int
   MimeType string
   QualityLabel string
   URL string
   Width int
}

func (f Format) Ext() (string, error) {
   exts, err := mime.ExtensionsByType(f.MimeType)
   if err != nil {
      return "", err
   }
   for _, ext := range exts {
      return ext, nil
   }
   return "", notPresent{f.MimeType}
}

func (f Format) Format(s fmt.State, verb rune) {
   if f.QualityLabel != "" {
      fmt.Fprint(s, "Quality:", f.QualityLabel)
   } else {
      fmt.Fprint(s, "Quality:", f.AudioQuality)
   }
   fmt.Fprint(s, " Bitrate:", f.AverageBitrate)
   fmt.Fprint(s, " Size:", f.ContentLength)
   fmt.Fprint(s, " Type:", f.MimeType)
   if verb == 'a' {
      fmt.Fprint(s, " URL:", f.URL)
   }
}

func (f Format) Write(dst io.Writer) error {
   req, err := http.NewRequest("GET", f.URL, nil)
   if err != nil {
      return err
   }
   LogLevel.Dump(req)
   var (
      begin = time.Now()
      content int64
   )
   for content < f.ContentLength {
      req.Header.Set(
         "Range", fmt.Sprintf("bytes=%v-%v", content, content+partLength-1),
      )
      end := time.Since(begin).Seconds()
      if end >= 1 {
         fmt.Print(format.Percent(content, f.ContentLength), "\t")
         fmt.Print(format.LabelSize(content), "\t")
         fmt.Println(format.LabelRate(float64(content)/end))
      }
      // this sometimes redirects, so cannot use http.Transport
      res, err := new(http.Client).Do(req)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      if _, err := io.Copy(dst, res.Body); err != nil {
         return err
      }
      content += partLength
   }
   return nil
}

// We cannot use bitrate to sort, as you end up with different heights:
//
// ID          | 480    | 720 low | 720 high | 1080
// ------------|--------|---------|----------|-----
// 7WTEB7Qbt4U | 285106 | 286687  | 513601   | 513675
// RPjE9riEhtA | 584072 | 1169166 | 1693812  | 2151670
type Height struct {
   StreamingData
   Target int
}

func (h Height) Less(i, j int) bool {
   return h.distance(i) < h.distance(j)
}

func (h Height) distance(i int) int {
   diff := h.AdaptiveFormats[i].Height - h.Target
   if diff >= 0 {
      return diff
   }
   return -diff
}

type notPresent struct {
   value string
}

func (n notPresent) Error() string {
   return fmt.Sprintf("%q is not present", n.value)
}
