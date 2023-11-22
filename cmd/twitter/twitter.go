package main

import (
   "fmt"
   "github.com/89z/format/hls"
   "github.com/89z/mech/twitter"
   "net/http"
   "os"
)

func doStatus(id, bitrate int64, info bool) error {
   guest, err := twitter.NewGuest()
   if err != nil {
      return err
   }
   stat, err := guest.Status(id)
   if err != nil {
      return err
   }
   if info {
      fmt.Println(stat)
   } else {
      for _, media := range stat.Extended_Entities.Media {
         for _, variant := range media.Variants() {
            if variant.Bitrate == bitrate {
               fmt.Println("GET", variant.URL)
               res, err := http.Get(variant.URL)
               if err != nil {
                  return err
               }
               defer res.Body.Close()
               ext, err := variant.Ext()
               if err != nil {
                  return err
               }
               dst, err := os.Create(stat.Base(id) + ext)
               if err != nil {
                  return err
               }
               defer dst.Close()
               if _, err := dst.ReadFrom(res.Body); err != nil {
                  return err
               }
            }
         }
      }
   }
   return nil
}

func doSpace(id string, info bool) error {
   guest, err := twitter.NewGuest()
   if err != nil {
      return err
   }
   space, err := guest.AudioSpace(id)
   if err != nil {
      return err
   }
   if info {
      fmt.Println(space)
   } else {
      source, err := guest.Source(space)
      if err != nil {
         return err
      }
      fmt.Println("GET", source.Location)
      res, err := http.Get(source.Location)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      seg, err := hls.NewSegment(res.Request.URL, res.Body)
      if err != nil {
         return err
      }
      file, err := os.Create(space.Base() + seg.Ext())
      if err != nil {
         return err
      }
      defer file.Close()
      for i, info := range seg.Info {
         fmt.Print(seg.Progress(i))
         res, err := http.Get(info.URI.String())
         if err != nil {
            return err
         }
         if _, err := file.ReadFrom(res.Body); err != nil {
            return err
         }
         if err := res.Body.Close(); err != nil {
            return err
         }
      }
   }
   return nil
}
