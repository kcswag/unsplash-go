package search

import (
	"business/application/unsplash/http"
	"business/application/unsplash/infrastructure/constants"
	"business/application/unsplash/proto"
)

type searchOrderBy string
var SOrderBy = map[string]searchOrderBy{
	"LATEST" : "latest",
	"RELEVANT" : "relevant",
}


type BasicParams struct{
	Query string
	Page string
	PerPage string
}

type PhotosParams struct {
	Query string
	Page int64
	PerPage int64
	OrderBy searchOrderBy
	Collections string
	ContentFilter constants.ContentFilter
	Color constants.Color
	Orientation constants.Orientation
}

type CollectionsParams struct {
	Query string
	Page int64
	PerPage int64
}

type UsersParams struct {
	Query string
	Page int64
	PerPage int64
}

func (p PhotosParams) Search() (*proto.PhotosSearch, *proto.ErrorResponse) {
	uri := "/search/photos"
	res := new(proto.PhotosSearch)
	errs := http.DoGet(uri,p,res)
	if errs != nil{
		return nil, errs
	}
	return res, nil
}

func (p CollectionsParams) Search() (*proto.CollectionsSearch, *proto.ErrorResponse) {
	uri := "/search/collections"
	res := new(proto.CollectionsSearch)
	errs := http.DoGet(uri,p,res)
	if errs != nil{
		return nil, errs
	}
	return res, nil
}

func (p UsersParams) Search() (*proto.UsersSearch, *proto.ErrorResponse) {
	uri := "/search/users"
	res := new(proto.UsersSearch)
	errs := http.DoGet(uri,p,res)
	if errs != nil{
		return nil, errs
	}
	return res, nil
}
