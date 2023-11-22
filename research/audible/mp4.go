package main

import (
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "time"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "d1jobzhhm62zby.cloudfront.net"
   req.URL.Path = "/bk_adbl_003303/2/signed/g1/bk_adbl_003303_22_64.mp4"
   req.URL.Scheme = "https"
   req.Header["Range"] = []string{"bytes=0-9999999"}
   req.Header["User-Agent"] = []string{"com.audible.playersdk.player/3.21.0 (Linux;Android 7.0) ExoPlayerLib/2.14.2"}
   val := make(url.Values)
   val["X-Amz-Date"] = []string{"20220217T044820Z"}
   val["X-Amz-Expires"] = []string{"86400"}
   val["X-Amz-Signature"] = []string{"8d59f224bcc663263b22578be1e3586e93cb710d697dfbba7c262489aa51bcc2"}
   val["X-Amz-SignedHeaders"] = []string{"host;user-agent"}
   val["id"] = []string{"8a3ac406-656e-444a-893e-3665dc9a0523"}
   req.URL.RawQuery = val.Encode()
   time.Sleep(time.Second)
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   buf, err := httputil.DumpResponse(res, false)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(buf)
}
