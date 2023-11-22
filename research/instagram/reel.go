package main

import (
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "i.instagram.com"
   req.URL.Path = "/api/v1/feed/reels_media/"
   val := make(url.Values)
   val["reel_ids"] = []string{"highlight:17939011231603196"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   req.Header["X-Ig-App-Id"] = []string{"936619743392459"}
   req.Header["Cookie"] = []string{sessionID}
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   buf, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(buf)
}
