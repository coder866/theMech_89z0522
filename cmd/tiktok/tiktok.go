package main

import (
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/tiktok"
   "net/http"
   "os"
)

func detail(awemeID int64, info bool) error {
   det, err := tiktok.NewDetail(awemeID)
   if err != nil {
      return err
   }
   if info {
      fmt.Println(det)
   } else {
      addr := det.URL()
      fmt.Println("GET", addr)
      res, err := http.Get(addr)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      ext, err := mech.Ext(res.Header)
      if err != nil {
         return err
      }
      file, err := os.Create(det.Base() + ext)
      if err != nil {
         return err
      }
      defer file.Close()
      if _, err := file.ReadFrom(res.Body); err != nil {
         return err
      }
   }
   return nil
}
