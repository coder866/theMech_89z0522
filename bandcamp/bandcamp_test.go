package bandcamp

import (
   "fmt"
   "testing"
   "time"
)

const artID = 3809045440

type tokenTest struct {
   typ, addr string
}

var tests = []tokenTest{
   {"a", "https://schnaussandmunk.bandcamp.com/album/passage-2"},
   {"i", "https://schnaussandmunk.bandcamp.com/music"},
   {"t", "https://schnaussandmunk.bandcamp.com/track/amaris-2"},
}

func TestImage(t *testing.T) {
   for _, img := range Images {
      addr := img.Format(artID)
      fmt.Println(addr)
   }
}

func TestItem(t *testing.T) {
   for _, test := range tests {
      tok, err := NewItem(test.addr)
      if err != nil {
         t.Fatal(err)
      }
      if tok.Item_Type != test.typ {
         t.Fatal(tok)
      }
      time.Sleep(99 * time.Millisecond)
   }
}
