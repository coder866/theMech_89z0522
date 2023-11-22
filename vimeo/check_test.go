package vimeo

import (
   "fmt"
   "testing"
)

const (
   id = 412573977
   password = "butter"
)

func TestCheck(t *testing.T) {
   LogLevel = 1
   check, err := Clip{ID: id}.Check(password)
   if err != nil {
      t.Fatal(err)
   }
   for _, pro := range check.Request.Files.Progressive {
      fmt.Println(pro)
   }
}
