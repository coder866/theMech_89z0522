package instagram

import (
   "fmt"
   "os"
   "testing"
   "time"
)

var usernames = []string{
   "karajewelll",
   "lokalist.id",
}

type testType struct {
   shortcode string
   paths []string
}

var tests = []testType{
   // image:0 video:1 DASH:1
   {"CLHoAQpCI2i", []string{
      "/v/t50.2886-16/147443688_171418268077432_595353563536612094_n.mp4",
   }},
   // image:1 video:0
   {"CZVEugIPkVn", []string{
      "/v/t51.2885-15/e35/272868603_955434698415740_8419357209460643893_n.jpg",
   }},
   // image:2 video:0
   {"CZAUQ_OvWZC", []string{
      "/v/t51.2885-15/e35/272178059_418057813408942_8652558621028999033_n.jpg",
      "/v/t51.2885-15/e35/272193572_305510171620132_8506371495778119983_n.jpg",
   }},
   // image:0 video:3 DASH:0
   {"BQ0eAlwhDrw", []string{
      "/v/t50.2886-16/16936668_174021943097999_6018573358768062464_n.mp4",
      "/v/t50.2886-16/16914924_575567635981228_3911260849125195776_n.mp4",
      "/v/t50.2886-16/16812535_590210974523279_620796230821216256_n.mp4",
   }},
   // image:2 video:1 DASH:1
   {"CUK-1wjqqsP", []string{
      "/v/t51.2885-15/e35/242545662_1278053282609020_5170310197887813120_n.jpg",
      "/v/t50.2886-16/242908146_4508662539195308_3750958489654012960_n.mp4",
      "/v/t50.2886-16/242707077_843418709708029_3826314158202589747_n.mp4",
      "/v/t51.2885-15/e35/242495948_255838523113734_4044316450944265352_n.jpg",
   }},
}

func TestUnderscoreMedia(t *testing.T) {
   cache, err := os.UserCacheDir()
   if err != nil {
      t.Fatal(err)
   }
   login, err := OpenLogin(cache + "/mech/instagram.json")
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      items, err := login.Items(test.shortcode)
      if err != nil {
         t.Fatal(err)
      }
      for _, item := range items {
         addrs, err := item.URLs()
         if err != nil {
            t.Fatal(err)
         }
         for _, addr := range addrs {
            fmt.Println(addr)
         }
      }
      fmt.Println()
      time.Sleep(time.Second)
   }
}

func TestUnderscoreUser(t *testing.T) {
   cache, err := os.UserCacheDir()
   if err != nil {
      t.Fatal(err)
   }
   login, err := OpenLogin(cache + "/mech/instagram.json")
   if err != nil {
      t.Fatal(err)
   }
   for _, username := range usernames {
      user, err := login.User(username)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(user)
      time.Sleep(time.Second)
   }
}
