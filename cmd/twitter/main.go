package main

import (
   "flag"
   "github.com/89z/mech/twitter"
)

func main() {
   // b
   var status int64
   flag.Int64Var(&status, "b", 0, "status ID")
   // c
   var space string
   flag.StringVar(&space, "c", "", "space ID")
   // f
   var bitrate int64
   flag.Int64Var(&bitrate, "f", 2_176_000, "status bitrate")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      twitter.LogLevel = 1
   }
   if status >= 1 {
      err := doStatus(status, bitrate, info)
      if err != nil {
         panic(err)
      }
   } else if space != "" {
      err := doSpace(space, info)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
