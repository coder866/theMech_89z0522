package youtube

import (
   "fmt"
   "net/http"
   "testing"
   "time"
)

const id = "UpNXI3_ctAc"

func TestImageFormat(t *testing.T) {
   for _, img := range Images {
      addr := img.Format(id)
      fmt.Println("HEAD", addr)
      res, err := http.Head(addr)
      if err != nil {
         t.Fatal(err)
      }
      if res.StatusCode != http.StatusOK {
         t.Fatal(res.Status)
      }
      time.Sleep(99 * time.Millisecond)
   }
}
