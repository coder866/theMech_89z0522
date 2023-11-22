package mech

import (
   "mime"
   "net/http"
   "strconv"
   "strings"
)

func Clean(path string) string {
   fn := func(r rune) rune {
      if strings.ContainsRune(`"*/:<>?\|`, r) {
         return -1
      }
      return r
   }
   return strings.Map(fn, path)
}

func Ext(head http.Header) (string, error) {
   typ := head.Get("Content-Type")
   exts, err := mime.ExtensionsByType(typ)
   if err != nil {
      return "", err
   }
   for _, ext := range exts {
      return ext, nil
   }
   return "", notPresent{typ}
}

type notPresent struct {
   value string
}

func (n notPresent) Error() string {
   return strconv.Quote(n.value) + " is not present"
}
