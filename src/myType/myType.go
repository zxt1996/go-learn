package myType

import (
	"fmt"
	"reflect"
	"strconv"
)

//强类型
func addition(x int, y int) int  {
	return x + y
}

func aboutString(s string) string  {
	return "hello " + s
}

func aboutType()  {
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

func variable()  {
	var s, t string = "one", "two"
	fmt.Println(s, t)

	//简短变量声明
	dog := "真的很狗"
	fmt.Println(dog)

	//块级作用域
	can := true
	if can {
		fmt.Println(dog)
	}
}

const change string = "不能改变的常量"

func MyType()  {
	aboutType()
	variable()

	fmt.Println(change)
}
