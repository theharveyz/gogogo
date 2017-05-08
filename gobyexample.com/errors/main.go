package main

import (
	"fmt"
)

type customError struct {
	code    uint16
	message string
}

// Error方法 无参数,返回string
func (ce *customError) Error() string {
	return fmt.Sprintf("%d - %s", ce.code, ce.message)
}

func newCustomErr(code uint16, msg string) error {
	return &customError{code, msg} // & 建议返回指针
}

func main() {
	fmt.Println(newCustomErr(500, "server error"))
}
