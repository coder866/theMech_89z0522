package main

import (
   "flag"
   "github.com/89z/mech/instagram"
)

func main() {
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // b
   var shortcode string
   flag.StringVar(&shortcode, "b", "", "shortcode")
   // h
   var auth bool
   flag.BoolVar(&auth, "h", false, "authentication")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   // p
   var password string
   flag.StringVar(&password, "p", "", "password")
   // u
   var username string
   flag.StringVar(&username, "u", "", "username")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      instagram.LogLevel = 1
   }
   if username != "" {
      err := saveLogin(username, password)
      if err != nil {
         panic(err)
      }
   } else if shortcode != "" || address != "" {
      if shortcode == "" {
         shortcode = instagram.Shortcode(address)
      }
      if auth {
         err := doItems(shortcode, info)
         if err != nil {
            panic(err)
         }
      } else {
         err := doGraph(shortcode, info)
         if err != nil {
            panic(err)
         }
      }
   } else {
      flag.Usage()
   }
}
