package youtube

import (
   "encoding/json"
   "net/http"
   "net/url"
   "os"
   "path/filepath"
   "strings"
)

func (x Exchange) Create(elem ...string) error {
   name := filepath.Join(elem...)
   err := os.MkdirAll(filepath.Dir(name), os.ModeDir)
   if err != nil {
      return err
   }
   os.Stdout.WriteString("Create " + name + "\n")
   file, err := os.Create(name)
   if err != nil {
      return err
   }
   defer file.Close()
   return json.NewEncoder(file).Encode(x)
}

func OpenExchange(elem ...string) (*Exchange, error) {
   file, err := os.Open(filepath.Join(elem...))
   if err != nil {
      return nil, err
   }
   defer file.Close()
   exc := new(Exchange)
   if err := json.NewDecoder(file).Decode(exc); err != nil {
      return nil, err
   }
   return exc, nil
}

const (
   // YouTube on TV
   clientID =
      "861556708454-d6dlm3lh05idd8npek18k6be8ba3oc68" +
      ".apps.googleusercontent.com"
   clientSecret = "SboVhoG9s0rNafixCSGGKXAT"
)

type Exchange struct {
   Access_Token string
   Error string
   Refresh_Token string
}

func (x Exchange) Header() http.Header {
   head := make(http.Header)
   head.Set("Authorization", "Bearer " + x.Access_Token)
   return head
}

func (x *Exchange) Refresh() error {
   val := url.Values{
      "client_id": {clientID},
      "client_secret": {clientSecret},
      "grant_type": {"refresh_token"},
      "refresh_token": {x.Refresh_Token},
   }
   res, err := http.PostForm("https://oauth2.googleapis.com/token", val)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(x)
}

type OAuth struct {
   Device_Code string
   User_Code string
   Verification_URL string
}

func NewOAuth() (*OAuth, error) {
   val := url.Values{
      "client_id": {clientID},
      "scope": {"https://www.googleapis.com/auth/youtube"},
   }
   res, err := http.PostForm("https://oauth2.googleapis.com/device/code", val)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   oau := new(OAuth)
   if err := json.NewDecoder(res.Body).Decode(oau); err != nil {
      return nil, err
   }
   return oau, nil
}

func (o OAuth) Exchange() (*Exchange, error) {
   val := url.Values{
      "client_id": {clientID},
      "client_secret": {clientSecret},
      "device_code": {o.Device_Code},
      "grant_type":  {"urn:ietf:params:oauth:grant-type:device_code"},
   }
   res, err := http.PostForm("https://oauth2.googleapis.com/token", val)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   exc := new(Exchange)
   if err := json.NewDecoder(res.Body).Decode(exc); err != nil {
      return nil, err
   }
   return exc, nil
}

func (o OAuth) String() string {
   var buf strings.Builder
   buf.WriteString("1. Go to\n")
   buf.WriteString(o.Verification_URL)
   buf.WriteString("\n\n2. Enter this code\n")
   buf.WriteString(o.User_Code)
   buf.WriteString("\n\n3. Press Enter to continue")
   return buf.String()
}
