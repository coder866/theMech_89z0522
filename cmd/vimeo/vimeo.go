package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/vimeo"
   "net/http"
   "net/url"
   "os"
   "path"
)

func doAuth(addr string, height int, info bool) error {
   return nil
}

func doAnon(addr string, height int, info bool) error {
   web, err := vimeo.NewJsonWeb()
   if err != nil {
      return err
   }
   clip, err := vimeo.NewClip(addr)
   if err != nil {
      return err
   }
   video, err := web.Video(clip)
   if err != nil {
      return err
   }
   if info {
      fmt.Println(video)
   } else {
      for _, down := range video.Download {
         if down.Height == height {
            fmt.Println("GET", down.Link)
            res, err := http.Get(down.Link)
            if err != nil {
               return err
            }
            defer res.Body.Close()
            addr, err := url.Parse(down.Link)
            if err != nil {
               return err
            }
            file, err := os.Create(path.Base(addr.Path))
            if err != nil {
               return err
            }
            defer file.Close()
            pro := format.NewProgress(res)
            if _, err := file.ReadFrom(pro); err != nil {
               return err
            }
         }
      }
   }
   return nil
}
