package tiktok

import (
   "fmt"
   "testing"
)

// tiktok.com/@kaimanoff/video/6896523341402737921
const id = 6896523341402737921

func TestDetail(t *testing.T) {
   det, err := NewDetail(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(det)
}
