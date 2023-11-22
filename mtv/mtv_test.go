package mtv

import (
   "fmt"
   "testing"
   "time"
)

var tests = []string{
   "https://www.mtv.com/episodes/scyb0g/aeon-flux-utopia-or-deuteranopia-season-1-ep-1",
   "https://www.mtv.com/video-clips/s5iqyc/mtv-cribs-dj-khaled",
}

func TestProperty(t *testing.T) {
   for _, test := range tests {
      prop, err := NewItem(test).Property()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", prop)
      time.Sleep(time.Second)
   }
}

func TestTopaz(t *testing.T) {
   prop, err := NewItem(tests[1]).Property()
   if err != nil {
      t.Fatal(err)
   }
   top, err := prop.Topaz()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", top)
}
