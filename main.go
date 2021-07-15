package main

import (
	"business/application/unsplash/application/auth"
	"business/application/unsplash/application/photo"
	"business/application/unsplash/infrastructure/constants"
	"business/application/unsplash/infrastructure/logger"
	"fmt"
	"reflect"
)

func main()  {
	p := new(photo.PhotosParams)
	p.Page = 10
	p.PerPage = 2
	callMethod(p,"ListPhotos")
}

func getAccessToken()  {
	p := new(auth.AccessTokenParams)
	p.ClientId = constants.ACCESS_KEY
	p.ClientSecret = constants.SECRET_KEY
	p.RedirectUri = "http://42.193.163.234:7777/redirect"
	p.Code = "ZtV6RwRF2007XVlNBsiWBa_Cngn-4NHox_sSuSDWJWI"
	p.GrantType = "authorization_code"

	res, errs := p.GetAccessToken()

	if errs != nil{
		logger.Error.Println(errs.Errors)
	}else{
		fmt.Println(res.AccessToken)
	}
}

func callMethod(params interface{}, methodName string)  {
	v := reflect.ValueOf(params)
	res := v.MethodByName(methodName).Call(nil)
	fmt.Println(res[0].Elem().Interface())
}
