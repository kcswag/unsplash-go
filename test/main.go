package main

import (
	"fmt"
	"reflect"
)

type proto interface {
	a()
}
type ins1 struct {
	aa string
}
type ins2 struct {
	bb string
}
func (i *ins1) a()  {}
func (i *ins2) a()  {}

type api interface {
	request() proto
}
type params1 struct{
	lk string
}
type params2 struct{
	lo string
}
func (p *params1) request() *ins1 {
	ins := new(ins1)
	ins.aa = "ssss"
	return ins
}
func (p *params2) request() *ins2 {
	ins := new(ins2)
	ins.bb = "ttttt"
	return ins
}


func main()  {
	p := new(params2)
	k := reflect.ValueOf(*p).Elem()
	fmt.Println(k.NumField())
}

func call(a api) proto  {
	t := reflect.TypeOf(a).String()
	fmt.Println(t)
	ins := new(ins1)
	ins.aa = "gkgkgkgkgk"
	return ins
}

