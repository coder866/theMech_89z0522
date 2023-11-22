package pandora

import (
   "fmt"
   "os"
   "testing"
)

const addr =
   "https://www.pandora.com/artist/the-black-dog/radio-scarecrow" +
   "/train-by-the-autobahn-part-1/TRddpp5JJ2hqnVV"

func TestAppLink(t *testing.T) {
   id, err := AppLink(addr)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%q\n", id)
}

func TestCreate(t *testing.T) {
   part, err := NewPartnerLogin()
   if err != nil {
      t.Fatal(err)
   }
   tLen := len(part.Result.PartnerAuthToken)
   if tLen != 34 {
      t.Fatal(tLen)
   }
   user, err := part.UserLogin("srpen6@gmail.com", password)
   if err != nil {
      t.Fatal(err)
   }
   if tLen := len(user.Result.UserAuthToken); tLen != 58 {
      t.Fatal(tLen)
   }
   cache, err := os.UserCacheDir()
   if err != nil {
      t.Fatal(err)
   }
   if err := user.Create(cache + "/mech/pandora.json"); err != nil {
      t.Fatal(err)
   }
}

