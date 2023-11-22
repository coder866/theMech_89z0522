package youtube

import (
   "encoding/base64"
   "github.com/89z/format/protobuf"
)

var Duration = map[string]protobuf.Uint64{
   "Under 4 minutes": 1,
   "4 - 20 minutes": 3,
   "Over 20 minutes": 2,
}

var Features = map[string]protobuf.Number{
   "360°": 15,
   "3D": 7,
   "4K": 14,
   "Creative Commons": 6,
   "HD": 4,
   "HDR": 25,
   "Live": 8,
   "Location": 23,
   "Purchased": 9,
   "Subtitles/CC": 5,
   "VR180": 26,
}

var SortBy = map[string]protobuf.Uint64{
   "Relevance": 0,
   "Upload date": 2,
   "View count": 3,
   "Rating": 1,
}

var Type = map[string]protobuf.Uint64{
   "Video": 1,
   "Channel": 2,
   "Playlist": 3,
   "Movie": 4,
}

var UploadDate = map[string]protobuf.Uint64{
   "Last hour": 1,
   "Today": 2,
   "This week": 3,
   "This month": 4,
   "This year": 5,
}

type Filter struct {
   protobuf.Message
}

func NewFilter() Filter {
   var filter Filter
   filter.Message = make(protobuf.Message)
   return filter
}

type Params struct {
   protobuf.Message
}

func NewParams() Params {
   var par Params
   par.Message = make(protobuf.Message)
   return par
}

func (p Params) Encode() string {
   buf := p.Marshal()
   return base64.StdEncoding.EncodeToString(buf)
}

func (f Filter) Features(num protobuf.Number) {
   f.Message[num] = protobuf.Uint64(1)
}

func (f Filter) Duration(val protobuf.Uint64) {
   f.Message[3] = val
}

func (f Filter) Type(val protobuf.Uint64) {
   f.Message[2] = val
}

func (f Filter) UploadDate(val protobuf.Uint64) {
   f.Message[1] = val
}

func (p Params) Filter(val Filter) {
   p.Message[2] = val.Message
}

func (p Params) SortBy(val protobuf.Uint64) {
   p.Message[1] = val
}
