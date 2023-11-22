package twitter

import (
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
   "net/url"
   "strings"
   "time"
)

const spacePersistedQuery = "lFpix9BgFDhAMjn9CrW6jQ"

func (g Guest) AudioSpace(id string) (*AudioSpace, error) {
   var str strings.Builder
   str.WriteString("https://twitter.com/i/api/graphql/")
   str.WriteString(spacePersistedQuery)
   str.WriteString("/AudioSpaceById")
   req, err := http.NewRequest("GET", str.String(), nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + bearer},
      "X-Guest-Token": {g.Guest_Token},
   }
   buf, err := json.Marshal(spaceRequest{ID: id})
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "variables=" + url.QueryEscape(string(buf))
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var space struct {
      Data struct {
         AudioSpace AudioSpace
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&space); err != nil {
      return nil, err
   }
   return &space.Data.AudioSpace, nil
}

type AudioSpace struct {
   Metadata struct {
      Media_Key string
      Title string
      Started_At int64
      Ended_At int64 `json:"ended_at,string"`
   }
   Participants struct {
      Admins []struct {
         Display_Name string
      }
   }
}

func (a AudioSpace) Admins() string {
   var buf strings.Builder
   for i, admin := range a.Participants.Admins {
      if i >= 1 {
         buf.WriteByte(',')
      }
      buf.WriteString(admin.Display_Name)
   }
   return buf.String()
}

func (a AudioSpace) Base() string {
   var buf strings.Builder
   buf.WriteString(a.Admins())
   buf.WriteByte('-')
   buf.WriteString(a.Metadata.Title)
   return mech.Clean(buf.String())
}

func (a AudioSpace) Duration() time.Duration {
   dur := a.Metadata.Ended_At - a.Metadata.Started_At
   return time.Duration(dur) * time.Millisecond
}

func (a AudioSpace) String() string {
   var buf strings.Builder
   buf.WriteString("Key: ")
   buf.WriteString(a.Metadata.Media_Key)
   buf.WriteString("\nTitle: ")
   buf.WriteString(a.Metadata.Title)
   buf.WriteString("\nStarted: ")
   buf.WriteString(a.Time().String())
   buf.WriteString("\nDuration: ")
   buf.WriteString(a.Duration().String())
   buf.WriteString("\nAdmins: ")
   buf.WriteString(a.Admins())
   return buf.String()
}

func (a AudioSpace) Time() time.Time {
   return time.UnixMilli(a.Metadata.Started_At)
}

type spaceRequest struct {
   ID string `json:"id"`
   IsMetatagsQuery bool `json:"isMetatagsQuery"`
   WithBirdwatchPivots bool `json:"withBirdwatchPivots"`
   WithDownvotePerspective bool `json:"withDownvotePerspective"`
   WithReactionsMetadata bool `json:"withReactionsMetadata"`
   WithReactionsPerspective bool `json:"withReactionsPerspective"`
   WithReplays bool `json:"withReplays"`
   WithScheduledSpaces bool `json:"withScheduledSpaces"`
   WithSuperFollowsTweetFields bool `json:"withSuperFollowsTweetFields"`
   WithSuperFollowsUserFields bool `json:"withSuperFollowsUserFields"`
}

type Source struct {
   Location string // Segment
}

func (g Guest) Source(space *AudioSpace) (*Source, error) {
   var str strings.Builder
   str.WriteString("https://twitter.com/i/api/1.1/live_video_stream/status/")
   str.WriteString(space.Metadata.Media_Key)
   req, err := http.NewRequest("GET", str.String(), nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + bearer},
      "X-Guest-Token": {g.Guest_Token},
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var video struct {
      Source Source
   }
   if err := json.NewDecoder(res.Body).Decode(&video); err != nil {
      return nil, err
   }
   return &video.Source, nil
}
