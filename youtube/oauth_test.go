package youtube

import (
   "fmt"
   "testing"
   "time"
)

func TestOAuth(t *testing.T) {
   oau, err := NewOAuth()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf(`1. Go to
%v

2. Enter this code
%v
`, oau.Verification_URL, oau.User_Code)
   for range [9]struct{}{} {
      time.Sleep(9 * time.Second)
      exc, err := oau.Exchange()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", exc)
      if exc.Access_Token != "" {
         break
      }
   }
}
