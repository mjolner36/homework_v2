package main

import (
	"fmt"
)

func printNumber(ptrToNumber interface{}) {
	if ptrToNumber == nil {
		fmt.Println("nil")
		return
	}

	ptr, ok := ptrToNumber.(*int)
	if !ok {
		fmt.Println("not an *int")
		return
	}

	if ptr == nil {
		fmt.Println("nil pointer")
		return
	}

	fmt.Println(*ptr)
}

func main() {
	v := 10
	printNumber(&v)
	u := 10
	printNumber(u)
	var pv *int
	printNumber(pv)
	pv = &v
	printNumber(pv)
}
