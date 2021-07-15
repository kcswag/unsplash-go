package http

import (
	"business/application/unsplash/infrastructure/constants"
	"business/application/unsplash/infrastructure/logger"
	"business/application/unsplash/infrastructure/util"
	"business/application/unsplash/proto"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)


var (
	host = "https://api.unsplash.com"
	accessKey = constants.ACCESS_KEY
	accessToken = constants.ACCESS_TOKEN
)

type Method string

const (
	GET Method = "GET"
	POST Method = "POST"
	PUT Method = "PUT"
	DELETE Method = "DELETE"
)


func Call(method Method, uri string, parameters interface{}) ([]byte, *proto.ErrorResponse) {
	var url string
	if strings.Contains(uri,"http"){
		url = uri
	}else{
		url = host+uri
	}
	req, err := http.NewRequest(string(method), url, nil)
	if err != nil {
		logger.Error.Println(err)
	}

	//req.Header.Add("Authorization","Client-ID "+accessKey)
	req.Header.Add("Authorization","Bearer "+accessToken)
	req.Header.Add("Content-Type", "application/json")

	if parameters != nil{
		setQuery(req, parameters)
	}

	client := &http.Client{}
	rsp, reqErr := client.Do(req)
	if reqErr != nil{
		logger.Error.Println(reqErr)
		errs := new(proto.ErrorResponse)
		errs.Errors = append(errs.Errors, reqErr.Error())
		return nil, errs
	}

	defer rsp.Body.Close()
	body, readErr := ioutil.ReadAll(rsp.Body)
	if readErr != nil{
		logger.Error.Println(readErr)
		return programError(readErr)
	}

	if rsp.StatusCode != 200 && rsp.StatusCode != 201 && rsp.StatusCode != 204{
		errs := new(proto.ErrorResponse)
		if err := json.Unmarshal(body, errs); err != nil{
			logger.Error.Println(err)
		}
		logger.Warning.Println(string(body))
		return body, errs
	}

	return body, nil
}

func Call2(method Method, uri string, parameters interface{}) ([]byte, *proto.ErrorResponse) {
	var req *http.Request
	var err error
	var url string
	if strings.Contains(uri,"http"){
		url = uri
	}else{
		url = host+uri
	}

	switch method {
		case GET:
			req, err = http.NewRequest(string(method), url, nil)
			if parameters != nil{
				setQuery(req, parameters)
			}
		case POST:
			if parameters != nil{
				data, err := json.Marshal(parameters)
				if err != nil{
					logger.Error.Println(err)
				}
				body := bytes.NewReader(data)
				req, err = http.NewRequest(string(method), url, body)
			}else{
				req, err = http.NewRequest(string(method), url, nil)
			}
	}

	if err != nil {
		logger.Error.Println(err)
	}
	//req.Header.Add("Authorization","Client-ID "+accessKey)
	req.Header.Add("Authorization","Bearer "+accessToken)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	rsp, reqErr := client.Do(req)
	if reqErr != nil{
		logger.Error.Println(reqErr)
		errs := new(proto.ErrorResponse)
		errs.Errors = append(errs.Errors, reqErr.Error())
		return nil, errs
	}

	defer rsp.Body.Close()
	body, readErr := ioutil.ReadAll(rsp.Body)
	//fmt.Println(string(body))
	if readErr != nil{
		logger.Error.Println(readErr)
		return programError(readErr)
	}

	if rsp.StatusCode != 200{
		errs := new(proto.ErrorResponse)
		if err := json.Unmarshal(body, errs); err != nil{
			logger.Error.Println(err)
		}
		logger.Warning.Println(string(body))
		return body, errs
	}

	return body, nil
}

func setQuery(req *http.Request, parameters interface{})  {
	q := req.URL.Query()
	pType := reflect.TypeOf(parameters)
	pValue := reflect.ValueOf(parameters)

	kind := pType.Kind()
	if kind == reflect.Struct {
		for k:=0; k < pType.NumField();k++{
			field := pType.Field(k).Name
			value := pValue.Field(k).String()
			if value != ""{
				q.Add(util.ToSnakeCase(field),util.ToSnakeCase(value))
			}
		}
	}else if kind == reflect.Ptr{
		t := pType.Elem()
		v := pValue.Elem()
		for k:=0; k < t.NumField();k++{
			field := t.Field(k)
			vField := v.Field(k)
			var value string
			if vField.Kind() == reflect.String{
				value = vField.String()
			}else if vField.Kind() == reflect.Int64{
				if vField.Int()!=0{
					value = strconv.FormatInt(vField.Int(),10)
				}
			}else if vField.Kind() == reflect.Float64{
				if vField.Float()!=0.00{
					value = strconv.FormatFloat(vField.Float(),'E',-1,64)
				}
			}else if vField.Kind() == reflect.Struct{
				for j:=0;j<vField.NumField();j++{
					sValue := vField.Field(j)
					if !sValue.IsZero(){
						fieldName := field.Name
						sField := field.Type.Field(j).Name
						q.Add(fmt.Sprintf("%s[%s]",strings.ToLower(fieldName),strings.ToLower(sField)), util.ToString(sValue.Interface()))
					}
				}
			}
			if value != ""{
				fmt.Println(value)
				//fmt.Println(util.ToSnakeCase(field)+":"+util.ToSnakeCase(value))
				q.Add(util.ToSnakeCase(field.Name),util.ToSnakeCase(value))
			}
		}
	}
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.RawQuery)
}

func DoGet(uri string, queryParams interface{}, res interface{}) *proto.ErrorResponse  {
	return doRequest(GET,uri,queryParams,res)
}

func DoPost(uri string, queryParams interface{}, res interface{}) *proto.ErrorResponse {
	return doRequest(POST,uri,queryParams,res)
}

func DoPut(uri string, queryParams interface{}, res interface{}) *proto.ErrorResponse {
	return doRequest(PUT,uri,queryParams,res)
}

func DoDelete(uri string, queryParams interface{}, res interface{}) *proto.ErrorResponse {
	return doRequest(DELETE,uri,queryParams,res)
}

func doRequest(method Method, uri string, params interface{}, res interface{}) *proto.ErrorResponse {
	rsp, errs := Call(method,uri,params)
	if errs == nil && res != nil{
		if err := json.Unmarshal(rsp,res); err != nil{
			logger.Error.Println(err)
		}
		return nil
	}
	return errs
}

func programError(err error) ([]byte, *proto.ErrorResponse) {
	errs := new(proto.ErrorResponse)
	errs.Errors = append(errs.Errors, err.Error())
	return nil, errs
}


