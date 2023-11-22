package twitter

import (
   "fmt"
   "testing"
)

const statusID = 1470124083547418624

func TestStatus(t *testing.T) {
   guest, err := NewGuest()
   if err != nil {
      t.Fatal(err)
   }
   stat, err := guest.Status(statusID)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", stat)
}
