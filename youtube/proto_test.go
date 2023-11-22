package youtube

import (
   "testing"
)

func TestProtoFilter(t *testing.T) {
   fil := NewFilter()
   fil.UploadDate(UploadDateLastHour)
   par := NewParams()
   par.Filter(fil)
   enc := par.Encode()
   if enc != "EgIIAQ==" {
      t.Fatal(enc)
   }
}

func TestProtoSort(t *testing.T) {
   par := NewParams()
   par.SortBy(SortByRating)
   enc := par.Encode()
   if enc != "CAE=" {
      t.Fatal(enc)
   }
}
