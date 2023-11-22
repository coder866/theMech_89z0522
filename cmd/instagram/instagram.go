package main

import (
   "fmt"
   "github.com/89z/mech/instagram"
   "net/http"
   "net/url"
   "os"
   "path/filepath"
   "time"
)

func doGraph(shortcode string, info bool) error {
   media, err := instagram.NewGraphMedia(shortcode)
   if err != nil {
      return err
   }
   if info {
      fmt.Println(media)
   } else {
      for _, addr := range media.URLs() {
         err := download(addr)
         if err != nil {
            return err
         }
         time.Sleep(99 * time.Millisecond)
      }
   }
   return nil
}

func doItems(shortcode string, info bool) error {
   cache, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   login, err := instagram.OpenLogin(cache + "/mech/instagram.json")
   if err != nil {
      return err
   }
   items, err := login.Items(shortcode)
   if err != nil {
      return err
   }
   for _, item := range items {
      if info {
         str, err := item.String()
         if err != nil {
            return err
         }
         fmt.Println(str)
      } else {
         addrs, err := item.URLs()
         if err != nil {
            return err
         }
         for _, addr := range addrs {
            err := download(addr)
            if err != nil {
               return err
            }
            time.Sleep(99 * time.Millisecond)
         }
      }
   }
   return nil
}

func saveLogin(username, password string) error {
   login, err := instagram.NewLogin(username, password)
   if err != nil {
      return err
   }
   cache, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   cache = filepath.Join(cache, "mech")
   os.Mkdir(cache, os.ModePerm)
   cache = filepath.Join(cache, "instagram.json")
   fmt.Println("Create", cache)
   return login.Create(cache)
}

func download(address string) error {
   fmt.Println("GET", address)
   res, err := http.Get(address)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   addr, err := url.Parse(address)
   if err != nil {
      return err
   }
   file, err := os.Create(filepath.Base(addr.Path))
   if err != nil {
      return err
   }
   defer file.Close()
   if _, err := file.ReadFrom(res.Body); err != nil {
      return err
   }
   return nil
}
