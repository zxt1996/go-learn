package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func addition(x int, y int) int  {
	return x + y
}

func aboutString(s string) string  {
	return "hello " + s
}

func main()  {
	fmt.Println(aboutString("jojo"))
	fmt.Println(addition(1, 2))

	var aboutBool bool
	fmt.Println(aboutBool)
	aboutBool = true
	fmt.Println(aboutBool)

	var f float32 = 0.111
	fmt.Println(f)

	var s string = "string"
	var i int = 1

	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(f))

	//类型转换
	var ss string = strconv.FormatBool(true)
	fmt.Println(ss)

	bb, err := strconv.ParseBool("true")
	fmt.Println(bb, err)
}
