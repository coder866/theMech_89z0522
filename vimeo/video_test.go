package vimeo

import (
   "fmt"
   "testing"
)

var clip = Clip{581039021, "9603038895"}

func TestVideo(t *testing.T) {
   web, err := NewJsonWeb()
   if err != nil {
      t.Fatal(err)
   }
   video, err := web.Video(&clip)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(video)
}

var videos = []string{
   "https://vimeo.com/477957994/2282452868",
   "https://player.vimeo.com/video/412573977?h=f7f2d6fcb7",
   "https://player.vimeo.com/video/412573977?unlisted_hash=f7f2d6fcb7",
   "https://vimeo.com/477957994?unlisted_hash=2282452868",
   "https://vimeo.com/66531465",
}

func TestClip(t *testing.T) {
   for _, video := range videos {
      clip, err := NewClip(video)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", clip)
   }
}
