package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/soundcloud"
   "time"
)

func main() {
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   // s
   var sleep time.Duration
   flag.DurationVar(&sleep, "s", time.Second, "sleep")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      soundcloud.LogLevel = 1
   }
   if address != "" {
      tracks, err := soundcloud.Resolve(address)
      if err != nil {
         panic(err)
      }
      for i, track := range tracks {
         if info {
            fmt.Printf("%+v\n", track)
         } else {
            if i >= 1 {
               time.Sleep(sleep)
            }
            err := download(track)
            if err != nil {
               panic(err)
            }
         }
      }
   } else {
      flag.Usage()
   }
}
