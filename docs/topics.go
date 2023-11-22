package main

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
   "net/url"
   "os"
)

const addr = "https://api.github.com/repos/89z/mech/topics"

var names = []string{
   "youtube",
   "instagram",
   "tiktok",
   "twitter",
   "soundcloud",
   "pandora",
   "vimeo",
   ///////////
   "paramount",
   "nbc",
   "mtv",
   "bandcamp",
}

func userinfo() (*url.Userinfo, error) {
   home, err := os.UserHomeDir()
   if err != nil {
      return nil, err
   }
   buf, err := os.ReadFile(home + "/.git-credentials")
   if err != nil {
      return nil, err
   }
   var addr url.URL
   if err := addr.UnmarshalBinary(bytes.TrimSpace(buf)); err != nil {
      return nil, err
   }
   return addr.User, nil
}

func main() {
   src := map[string][]string{"names": names}
   dst := new(bytes.Buffer)
   if err := json.NewEncoder(dst).Encode(src); err != nil {
      panic(err)
   }
   req, err := http.NewRequest("PUT", addr, dst)
   if err != nil {
      panic(err)
   }
   info, err := userinfo()
   if err != nil {
      panic(err)
   }
   password, ok := info.Password()
   if ok {
      req.SetBasicAuth(info.Username(), password)
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   fmt.Printf("%+v\n", res)
}
