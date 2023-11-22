package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/soundcloud"
   "net/http"
   "os"
)

func download(track soundcloud.Track) error {
   media, err := track.Progressive()
   if err != nil {
      return err
   }
   fmt.Println("GET", media.URL)
   res, err := http.Get(media.URL)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   ext, err := media.Ext()
   if err != nil {
      return err
   }
   file, err := os.Create(track.Base() + ext)
   if err != nil {
      return err
   }
   defer file.Close()
   pro := format.NewProgress(res)
   if _, err := file.ReadFrom(pro); err != nil {
      return err
   }
   return nil
}
