package main

import (
	"fmt"
	"reflect"
)

type Print func(string) string
type Print1 func(string) string

func pf(str string) string {
	fmt.Println(str)
	return "pf func"
}

func main() {
	pq := senFun()
	fmt.Println("type:", reflect.TypeOf(pq))
	_, ok := pq.(Print)
	fmt.Println(ok)
	pq = senFun()
	fmt.Println("type:", reflect.TypeOf(pq))
	_, ok = pq.(Print1)
	fmt.Println(ok)
}

func senFun() interface{} {
	var p Print
	p = pf
	return p
}
