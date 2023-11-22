package twitter

import (
   "encoding/json"
   "net/http"
   "strconv"
)

func (s Search) String() string {
   var buf []byte
   for key, val := range s.GlobalObjects.Tweets {
      if buf != nil {
         buf = append(buf, '\n')
      }
      buf = append(buf, "Tweet: "...)
      buf = strconv.AppendInt(buf, key, 10)
      for _, addr := range val.Entities.URLs {
         buf = append(buf, "\nURL: "...)
         buf = append(buf, addr.Expanded_URL...)
      }
   }
   return string(buf)
}

func (g Guest) Search(query string) (*Search, error) {
   return g.SearchCount(query, 0)
}

func (g Guest) SearchCount(query string, count int) (*Search, error) {
   req, err := http.NewRequest(
      "GET", "https://twitter.com/i/api/2/search/adaptive.json", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + bearer},
      "X-Guest-Token": {g.Guest_Token},
   }
   val := req.URL.Query()
   val.Set("q", query)
   if count >= 1 {
      val.Set("count", strconv.Itoa(count))
   }
   req.URL.RawQuery = val.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   sea := new(Search)
   if err := json.NewDecoder(res.Body).Decode(sea); err != nil {
      return nil, err
   }
   return sea, nil
}

type Search struct {
   GlobalObjects struct {
      Tweets map[int64]struct {
         Entities struct {
            URLs []struct {
               Expanded_URL string
            }
         }
      }
   }
}
