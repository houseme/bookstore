package main

import (
	"reflect"
)

// @Project: bookstore
// @Author: houseme
// @Description:
// @File: demo
// @Version: 1.0.0
// @Date: 2021/7/17 21:52
// @Package bookstore

type noop struct {
}

// Hello .
func (n *noop) Hello() {
	println("Noop")
}

// Demo .
func Demo() interface{ Hello() } {
	println("Demo")
	return new(noop)
}

func main() {
	Demo()
	Demo().Hello()

	var ptr = Demo()
	println(ptr)
	println(reflect.TypeOf(ptr).Elem().Name())
}
