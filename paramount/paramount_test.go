package paramount

import (
   "fmt"
   "testing"
   "time"
)

var guids = []string{
   // paramountplus.com/movies/building-star-trek/wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_
   "wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_",
   // paramountplus.com/shows/aeon-flux/video/IBmGAIrEofp2vQNvSEvO5MavtZy5GOAy/aeon-flux-isthmus-crypticus
   "IBmGAIrEofp2vQNvSEvO5MavtZy5GOAy",
}

func TestParamount(t *testing.T) {
   for _, guid := range guids {
      med, err := NewMedia(guid)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", med)
      time.Sleep(time.Second)
   }
}
