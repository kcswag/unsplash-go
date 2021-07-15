package collection

import (
	"business/application/unsplash/http"
	"business/application/unsplash/infrastructure/constants"
	"business/application/unsplash/proto"
	"fmt"
)

type ListParams struct {
	Page int64
	PerPage int64
}

func (p *ListParams) ListCollections() (*[]proto.Collection, *proto.ErrorResponse)  {
	uri := "/collections"
	collections := make([]proto.Collection,0)
	errs := http.DoGet(uri,p, &collections)
	return &collections,errs
}

type Id string

func (id Id) Get() (*proto.Collection, *proto.ErrorResponse)  {
	uri := fmt.Sprintf("/collections/%s",id)
	collection := new(proto.Collection)
	errs := http.DoGet(uri,id,collection)
	return collection,errs
}

type CollectionsPhotosParams struct {
	Id string  //required
	Page int64
	PerPage int64
	Orientation constants.Orientation
}

func (p *CollectionsPhotosParams) GetCollectionsPhotos() (*[]proto.Photo, *proto.ErrorResponse)  {
	uri := fmt.Sprintf("/collections/%s/photos", p.Id)
	photos := make([]proto.Photo,0)
	errs := http.DoGet(uri,p,&photos)
	return &photos,errs
}

func (id Id) ListRelatedCollections() (*[]proto.Collection, *proto.ErrorResponse)  {
	uri := fmt.Sprintf("/collections/%s/related", id)
	collections := make([]proto.Collection,0)
	errs := http.DoGet(uri,id,&collections)
	return &collections,errs
}

type CreateCollectionParams struct {
	Title string   //required
	Description string //optional
	Private constants.Bool
}

func (p *CreateCollectionParams) CreateCollection() (*proto.Collection,*proto.ErrorResponse) {
	uri := "/collections?title=kcdope"
	collection := new(proto.Collection)
	errs := http.DoPost(uri,p,collection)
	return collection,errs
}

type UpdateCollectionParams struct {
	Id string   //required
	Title string   //optional
	Description string  //optional
	Private string  //optional
}

func (p *UpdateCollectionParams) UpdateCollection() (*proto.Collection, *proto.ErrorResponse) {
	uri := fmt.Sprintf("/collections/%s", p.Id)
	collection := new(proto.Collection)
	errs := http.DoPut(uri,p,collection)
	return collection,errs
}

func (id Id) DeleteCollection() *proto.ErrorResponse {
	uri := fmt.Sprintf("/collections/%s", id)
	return http.DoDelete(uri,id,nil)
}

type PhotoOperationParams struct{
	CollectionId string `json:"collection_id"`   //required
	PhotoId string `json:"photo_id"`  //required
}

func (p *PhotoOperationParams) AddPhoto() (*proto.PhotoOperationOfCollection, *proto.ErrorResponse) {
	uri := fmt.Sprintf("/collections/%s/add", p.CollectionId)
	res := new(proto.PhotoOperationOfCollection)
	errs := http.DoPost(uri,p,res)
	return res, errs
}

func (p *PhotoOperationParams) RemovePhoto() (*proto.PhotoOperationOfCollection, *proto.ErrorResponse)  {
	uri := fmt.Sprintf("/collections/%s/remove",p.CollectionId)
	res := new(proto.PhotoOperationOfCollection)
	errs := http.DoPost(uri,p,res)
	return res, errs
}