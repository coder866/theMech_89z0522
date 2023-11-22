package soundcloud

import (
   "fmt"
   "testing"
   "time"
)

type itemType struct {
   id int64
   addr string
}

var items = []itemType{
   {936653761, "https://soundcloud.com/kino-scmusic/mqymd53jtwag"},
   {692707328, "https://soundcloud.com/kino-scmusic"},
}

func TestResolve(t *testing.T) {
   for _, item := range items {
      tracks, err := Resolve(item.addr)
      if err != nil {
         t.Fatal(err)
      }
      for _, track := range tracks {
         fmt.Printf("%+v\n", track)
      }
      time.Sleep(time.Second)
   }
}

func TestTrack(t *testing.T) {
   track, err := NewTrack(items[0].id)
   if err != nil {
      t.Fatal(err)
   }
   media, err := track.Progressive()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", media)
}
