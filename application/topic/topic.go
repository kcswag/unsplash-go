package topic

import (
	"business/application/unsplash/http"
	"business/application/unsplash/infrastructure/constants"
	"business/application/unsplash/infrastructure/logger"
	"business/application/unsplash/proto"
	"encoding/json"
	"fmt"
)

type listTopicsOrderBy string
var LTOrderBy = map[string]listTopicsOrderBy{
	"FEATURED":"featured",
	"LATEST":"latest",
	"OLDEST":"oldest",
	"POSITION":"position",
}

type topicsPhotosOrderBy string
var TPOrderBy = map[string]topicsPhotosOrderBy{
	"LATEST":"latest",
	"OLDEST":"oldest",
	"POPULAR":"popular",
}

type Topic string

type ListTopicsParams struct{
	Ids string `json:"ids"`
	Page int64 `json:"page"`
	PerPage int64 `json:"per_page"`
	OrderBy listTopicsOrderBy `json:"order_by"`
}

type PhotosParams struct {
	IdOrSlug string `json:"id_or_slug"`
	Page int64
	PerPage int64
	Orientation constants.Orientation
}

type TopicsPhotosParams struct {
	IdOrSlug string `json:"id_or_slug"`
	Page int64  //optional
	PerPage int64  //optional
	Orientation constants.Orientation  //optional
	OrderBy topicsPhotosOrderBy  //optional
}

func (t Topic) Get() (*proto.Topic, *proto.ErrorResponse)  {
	uri := fmt.Sprintf("/topics/%s", t)
	rsp, errs := http.Call(http.GET,uri,nil)

	topicData := new(proto.Topic)
	if errs == nil {
		if err := json.Unmarshal(rsp, topicData); err != nil{
			logger.Error.Println(err)
		}
		return topicData,nil
	}
	return nil,errs
}

func (p *ListTopicsParams) ListTopics() (*[]proto.Topic, *proto.ErrorResponse)  {
	uri := "/topics"
	rsp, errs := http.Call(http.GET, uri, p)

	if errs == nil{
		res := make([]proto.Topic,0)
		if err := json.Unmarshal(rsp, &res); err != nil{
			logger.Error.Println(err)
		}
		return &res, nil
	}
	return nil, errs
}

func (p *TopicsPhotosParams) GetTopicsPhotos() (*[]proto.Photo, *proto.ErrorResponse)  {
	uri := fmt.Sprintf("/topics/%s/photos", p.IdOrSlug)
	rsp, errs := http.Call(http.GET, uri, p)

	if errs == nil{
		res := make([]proto.Photo,0)
		if err := json.Unmarshal(rsp, &res); err != nil{
			logger.Error.Println(err)
		}
		return &res, nil
	}
	return nil, errs
}
