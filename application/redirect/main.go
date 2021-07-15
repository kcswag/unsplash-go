package main

import (
	"business/application/unsplash/infrastructure/logger"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {
	http.HandleFunc("/redirect", handle)

	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		logger.Error.Println(err)
	}

}

func handle(w http.ResponseWriter, r *http.Request)  {
	res, err := ioutil.ReadAll(r.Body)
	if err != nil{
		logger.Error.Println(err)
	}
	fmt.Println(string(res))
}
