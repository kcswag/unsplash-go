package util

import (
	"regexp"
	"strconv"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake  = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func ToString(i interface{}) string {
	var res string
	switch i.(type) {
	case float64:
		res = strconv.FormatFloat(i.(float64), 'E', -1, 64)
	case string:
		res = i.(string)
	case int64:
		res = strconv.FormatInt(i.(int64),10)
	case int:
		res = strconv.Itoa(i.(int))
	}
	return res
}
