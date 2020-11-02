package main

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

//指针
func pointer()  {
	s := "JOJO"
	//在变量名前加上&可获取变量在内存中的地址
	fmt.Println(&s)
}

//加上星号意味着参数的类型为指向整数的指针
func showMemoryAddress(x *int)  {
	fmt.Println(x)

	//要使用指针指向的变量的值，可在指针变量前加上星号
	fmt.Println(*x)
}

const change string = "不能改变的常量"

func main()  {
	aboutType()
	variable()
	pointer()

	i := 1
	showMemoryAddress(&i)

	fmt.Println(change)
}
