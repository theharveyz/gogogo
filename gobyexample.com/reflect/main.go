package main

import (
	"fmt"
	"reflect"
)

// Stuff this is doc
type Stuff struct {
	Erp   string `erp`
	Email string `email`
	_pri  string `pri:"test1" email:"test2"`
}

func main() {
	s := &Stuff{"zhanghuawei7", "zhanghuawei@jd.com", "pri"} // pointer:
	fmt.Println("stuff:", s)
	v := reflect.ValueOf(s)
	fmt.Println(v)

	rv := reflect.TypeOf(s)
	// fmt.Println(rv.Key())
	fmt.Println(rv.Kind())                // ptr
	fmt.Println(rv.Kind() == reflect.Ptr) // reflect value and reflect.ptr

	/**
	 * reflect.Struct
	 */

	fmt.Println(rv.Elem()) // main.Stuff

	// field & Field tag
	elms := rv.Elem()
	for i := 0; i < elms.NumField(); i++ {
		fmt.Println(elms.Field(i))                  // like: {Erp  string erp 0 [0] false}
		fmt.Println(elms.Field(i).Tag)              // tag
		fmt.Println(elms.Field(i).Tag.Get("email")) // tag group => test2
	}

}
