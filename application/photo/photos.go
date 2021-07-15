package photo

import (
	"business/application/unsplash/http"
	"business/application/unsplash/infrastructure/constants"
	"business/application/unsplash/proto"
	"fmt"
)

type orderBy string
var OrderBy = map[string]orderBy{
	"LATEST":"latest",
	"OLDEST":"oldest",
	"POPULAR":"popular",
}

type PhotosParams struct {
	Page int64 //optional, default 1
	PerPage int64 //optional, default 10
	OrderBy orderBy //optional, default latest
}

func (p *PhotosParams) ListPhotos() (*[]proto.Photo, *proto.ErrorResponse) {
	uri := "/photos"
	photos := make([]proto.Photo,0)
	errs := http.DoGet(uri,p,&photos)
	return &photos, errs
}

type Id string

func (id Id) Get() (*proto.Photo, *proto.ErrorResponse)  {
	uri := fmt.Sprintf("/photos/%s", id)
	photo := new(proto.Photo)
	errs := http.DoGet(uri,id,photo)
	return photo, errs
}


type RandomPhotoParams struct {
	Collections string
	Topics string
	Username string
	Query string
	Orientation constants.Orientation
	ContentFilter constants.ContentFilter `json:"content_filter"`
	Count int64
}

func (p *RandomPhotoParams) Get() (*proto.Photo, *proto.ErrorResponse)  {
	uri := "/photos/random"
	photo := new(proto.Photo)
	errs := http.DoGet(uri,p, photo)
	return photo, errs
}

type StatisticsParams struct {
	Id string  //required
	Resolution string  //optional, default "days"
	Quantity int64 //optional, default 30
}


func (p *StatisticsParams) Get() (*proto.PhotoStatistics, *proto.ErrorResponse)  {
	uri := fmt.Sprintf("/photos/%s/statistics", p.Id)
	res := new(proto.PhotoStatistics)
	errs := http.DoGet(uri,p,res)
	return res, errs
}

func (id Id) LikePhoto() (*proto.PhotoLike, *proto.ErrorResponse)  {
	uri := fmt.Sprintf("/photos/%s/like", id)
	res := new(proto.PhotoLike)
	errs := http.DoPost(uri,id,res)
	return res, errs
}

type PhotoUpdateParams struct {
	Id            string            `json:"id,omitempty"`
	Description   string            `json:"description,omitempty"`
	ShowOnProfile string            `json:"show_on_profile,omitempty"`
	Tags          string            `json:"tags,omitempty"`
	Location      Location            `json:"latitude,omitempty"`
	Exif          Exif `json:"exif,omitempty"`
}

type Location struct {
	Latitude float64
	Longitude float64
	Name string
	City string
	Country string
}

type Exif struct {
	Make            string `json:"make,omitempty"`
	Model           string `json:"model,omitempty"`
	ExposureTime    string `json:"exposure_time,omitempty"`
	ApertureValue   string `json:"aperture_value,omitempty"`
	FocalLength     string `json:"focal_length,omitempty"`
	IsoSpeedRatings string `json:"iso_speed_ratings,omitempty"`
}

func (p *PhotoUpdateParams) UpdatePhoto() (*proto.Photo, *proto.ErrorResponse) {
	uri := fmt.Sprintf("/photos/%s", p.Id)
	res := new(proto.Photo)
	errs := http.DoPut(uri,p,res)
	return res, errs
}