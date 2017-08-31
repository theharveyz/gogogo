package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var obj *Obj

type Obj struct {
	Test string
}

func getSingleton() *Obj {
	if obj == nil {
		once.Do(func() {
			obj = &Obj{}
		})
	}
	return obj
}

func main() {
	o := getSingleton()
	fmt.Println(o)
}
