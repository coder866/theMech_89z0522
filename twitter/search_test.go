package twitter

import (
   "fmt"
   "testing"
)

func TestSearch(t *testing.T) {
   guest, err := NewGuest()
   if err != nil {
      t.Fatal(err)
   }
   sea, err := guest.Search("filter:spaces")
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", sea)
}
