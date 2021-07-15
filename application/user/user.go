package user

import (
	"business/application/unsplash/http"
	"business/application/unsplash/infrastructure/constants"
	"business/application/unsplash/proto"
	"fmt"
)

type UsernameParams string

func (p UsernameParams) GetProfile() (*proto.User, *proto.ErrorResponse)  {
	uri := fmt.Sprintf("/users/%s", p)
	user := new(proto.User)
	errs := http.DoGet(uri,p,user)
	if errs != nil{
		return nil, errs
	}
	return user,nil
}

func (p UsernameParams) GetPortfolioLink() (*proto.PortfolioLink, *proto.ErrorResponse)  {
	uri := fmt.Sprintf("/users/%s/portfolio", p)
	res := new(proto.PortfolioLink)
	errs := http.DoGet(uri,p,res)
	if errs != nil {
		return nil, errs
	}
	return res, nil
}

type usersPhotosOrderBy string
var UsersPhotosOrderBy = map[string]usersPhotosOrderBy{
	"LATEST":"latest",
	"OLDEST":"oldest",
	"POPULAR":"popular",
	"VIEWS":"views",
	"DOWNLOADS":"downloads",
}


type ListUsersPhotosParams struct {
	Username string  //required
	Page string  //optional, default 1
	PerPage string  //optional, default 10
	OrderBy usersPhotosOrderBy //optional
	Stats constants.Bool //optional
	Resolution string
	Quantity int
	Orientation constants.Orientation
}

func (p *ListUsersPhotosParams) ListUsersPhotos() (*[]proto.Photo, *proto.ErrorResponse)  {
	uri := fmt.Sprintf("/users/%s/photos", p.Username)
	photos := make([]proto.Photo,0)
	errs := http.DoGet(uri,p,&photos)
	if errs != nil{
		return nil, errs
	}
	return &photos, nil
}

type usersLikedPhotosOrderBy string
var UsersLikedPhotosOrderBy = map[string]usersLikedPhotosOrderBy{
	"LATEST":"latest",
	"OLDEST":"oldest",
	"POPULAR":"popular",
}
type ListUsersLikedPhotosParams struct {
	Username string
	Page int64
	PerPage int64
	OrderBy usersLikedPhotosOrderBy
	Orientation constants.Orientation
}

func (p *ListUsersLikedPhotosParams) ListUsersLikedPhotos() (*[]proto.Photo, *proto.ErrorResponse)  {
	uri := fmt.Sprintf("/users/%s/likes",p.Username)
	photos := make([]proto.Photo,0)
	errs := http.DoGet(uri,p,&photos)
	if errs != nil{
		return nil, errs
	}
	return &photos,nil
}

type ListUsersCollectionsParams struct {
	Username string
	Page int64
	PerPage int64
}

func (p *ListUsersCollectionsParams) ListUsersCollections() (*[]proto.Collection, *proto.ErrorResponse)  {
	uri := fmt.Sprintf("/users/%s/collections", p.Username)
	collections := make([]proto.Collection, 0)
	errs := http.DoGet(uri,p,&collections)
	if errs != nil{
		return nil, errs
	}
	return &collections, nil
}

type UsersStatisticsParams struct {
	Username string
	Resolution string
	Quantity int64
}

func (p *UsersStatisticsParams) GetUsersStatistics() (*proto.UserStatistics, *proto.ErrorResponse)  {
	uri := fmt.Sprintf("/users/%s/statistics",p.Username)
	res := new(proto.UserStatistics)
	errs := http.DoGet(uri,p,res)
	if errs != nil{
		return nil, errs
	}
	return res,nil
}
