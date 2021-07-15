package auth

import (
	"business/application/unsplash/http"
	"business/application/unsplash/infrastructure/constants"
	"business/application/unsplash/infrastructure/logger"
	"business/application/unsplash/proto"
	"fmt"
	"io/ioutil"
	http2 "net/http"
)

var (
	oAuthUrl = "https://unsplash.com/oauth"
)

type AuthParams struct {
	ClientId string
	RedirectUri string
	ResponseType string
	Scope string
}

func Authorize()  {
	url := "https://unsplash.com/oauth/authorize"
	req, err := http2.NewRequest(http2.MethodPost,url, nil)
	if err != nil{
		panic(err)
	}
	fmt.Print(fmt.Sprintf("%v+", req))
	req.Header.Add("Cache-Control", "must-revalidate, no-store")
	req.Header.Add("Content-Type", " text/html;charset=UTF-8")
	req.Header.Add("Location", url)
	q := req.URL.Query()
	q.Add("client_id",constants.ACCESS_KEY)
	q.Add("redirect_uri","http://42.193.163.234:7777/redirect")
	q.Add("response_type","code")
	q.Add("scope","public+read_user+write_user+read_photos+write_photos+write_likes+write_followers+read_collections+write_collections")
	req.URL.RawQuery = q.Encode()

	client := &http2.Client{}
	rsp, reqErr := client.Do(req)
	if reqErr != nil{
		logger.Error.Println(reqErr)
		errs := new(proto.ErrorResponse)
		errs.Errors = append(errs.Errors, reqErr.Error())
	}

	body, readErr := ioutil.ReadAll(rsp.Body)
	if readErr != nil{
		panic(readErr)
	}

	fmt.Println(string(body))
}

type AccessTokenParams struct {
	ClientId string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectUri string `json:"redirect_uri"`
	Code string `json:"code"`
	GrantType string `json:"grant_type"`

}

func (p *AccessTokenParams) GetAccessToken() (*proto.Auth,*proto.ErrorResponse)  {
	uri := "https://unsplash.com/oauth/token"
	res := new(proto.Auth)
	errs := http.DoPost(uri,p,res)
	return res,errs
}


