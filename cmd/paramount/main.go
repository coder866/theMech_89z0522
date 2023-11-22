package main

import (
   "flag"
   "github.com/89z/mech/paramount"
)

func main() {
   // b
   var guid string
   flag.StringVar(&guid, "b", "", "GUID")
   // f
   var bandwidth int
   flag.IntVar(&bandwidth, "f", 3_063_000, "target bandwidth")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      paramount.LogLevel = 1
   }
   if guid != "" {
      err := doManifest(guid, bandwidth, info)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
