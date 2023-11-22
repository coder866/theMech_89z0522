package main

import (
   "fmt"
   "github.com/89z/format/hls"
   "github.com/89z/mech/mtv"
   "net/http"
   "net/url"
   "os"
   "sort"
)

func doManifest(addr string, bandwidth int, info bool) error {
   prop, err := mtv.NewItem(addr).Property()
   if err != nil {
      return err
   }
   top, err := prop.Topaz()
   if err != nil {
      return err
   }
   fmt.Println("GET", top.StitchedStream.Source)
   res, err := http.Get(top.StitchedStream.Source)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   mas, err := hls.NewMaster(res.Request.URL, res.Body)
   if err != nil {
      return err
   }
   if info {
      fmt.Println(prop.Data.Item)
      for _, str := range mas.Stream {
         fmt.Println(str)
      }
   } else {
      sort.Sort(hls.Bandwidth{mas, bandwidth})
      video := mas.Stream[0]
      err := download(prop, video.URI, fmt.Sprint(video))
      if err != nil {
         return err
      }
      return download(prop, mas.GetMedia(video).URI, "")
   }
   return nil
}

func download(prop *mtv.Property, addr *url.URL, stream string) error {
   seg, err := newSegment(addr)
   if err != nil {
      return err
   }
   fmt.Println("GET", seg.Key.URI)
   res, err := http.Get(seg.Key.URI.String())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   dec, err := hls.NewDecrypter(res.Body)
   if err != nil {
      return err
   }
   file, err := os.Create(prop.Base() + seg.Ext())
   if err != nil {
      return err
   }
   defer file.Close()
   if stream != "" {
      fmt.Println(stream)
   }
   for i, info := range seg.Info {
      fmt.Print(seg.Progress(i))
      res, err := http.Get(info.URI.String())
      if err != nil {
         return err
      }
      if _, err := dec.Copy(file, res.Body); err != nil {
         return err
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
   }
   return nil
}

func newSegment(addr *url.URL) (*hls.Segment, error) {
   fmt.Println("GET", addr)
   res, err := http.Get(addr.String())
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return hls.NewSegment(res.Request.URL, res.Body)
}
