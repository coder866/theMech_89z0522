package main

import (
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "time"
)

// authorization_code = parsed_url["openid.oa2.authorization_code"][0]

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "www.amazon.com"
   req.URL.Path = "/ap/signin"
   req.URL.Scheme = "https"
   val := make(url.Values)
   val["openid.assoc_handle"] = []string{"amzn_audible_android_aui_us"}
   val["openid.claimed_id"] = []string{"http://specs.openid.net/auth/2.0/identifier_select"}
   val["openid.identity"] = []string{"http://specs.openid.net/auth/2.0/identifier_select"}
   val["openid.mode"] = []string{"checkid_setup"}
   val["openid.ns"] = []string{"http://specs.openid.net/auth/2.0"}
   val["openid.ns.oa2"] = []string{"http://www.amazon.com/ap/ext/oauth/2"}
   val["openid.oa2.response_type"] = []string{"code"}
   val["openid.oa2.scope"] = []string{"device_auth_access"}
   val["openid.oa2.client_id"] = []string{"device:3738656232643031306334623466323238346237234131304b49535032475746304534"}
   val["openid.oa2.code_challenge"] = []string{"FqnF5AR7EuNjawwfQ2f757HcSMrEej9V3GqSsyzWS9Q"}
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
   os.Stderr.Write(buf)
   os.Stdout.ReadFrom(res.Body)
}
