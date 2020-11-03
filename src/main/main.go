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

//返回多个值
func moreNum() (int, string)  {
	one := 2
	two := "jojo"

	return one, two
}

//不定参数
func forMore(number... int)  {
	fmt.Println(number)
}

//具名返回值
//在终止语句return前给具名变量进行了赋值。使用具名返回值时，无须显式地返回相应的变量
func knowName() (x, y string) {
	x = "jojo"
	y = "ok"

	return
}

//将函数做为值传递
func anotherFunction(f func() string) string  {
	return f()
}

//控制流程
func controlFlow()  {
	can := true

	//if语句
	if can {
		fmt.Println("可以")
	}

	two := 12

	if two < 10 {
		fmt.Println("小于10")
	} else if two < 12 {
		fmt.Println("小于12")
	} else {
		fmt.Println("其他")
	}

	var all int
	for i := 0; i < 10; i++ {
		all += i
	}

	fmt.Println(all)

	//遍历数组
	numbers := []int {1, 2, 3}

	for index, value := range numbers{
		fmt.Println("numbers[", index, "]: ", value)
	}
}

//defer语句
//在函数返回前执行另一个函数
//通常用于执行清理操作或确保操作（如网络调用）完成后再执行另一个函数
func aboutDe()  {
	fmt.Println("start")

	defer fmt.Println("one")
	defer fmt.Println("two")

	fmt.Println("end")
}

//数组、切片和映射
func gather()  {
	var arr [2] int

	arr[0] = 1
	arr[1] = 2

	fmt.Println(arr)

	//内置函数make创建一个切片，其中第一个参数为数据类型，而第二个参数为长度
	var cheeses = make([] int, 3)
	for i := 0;i < 3;i++ {
		cheeses[i] = i
	}

	fmt.Println(cheeses)

	//往切片中添加元素
	cheeses = append(cheeses, 2, 3)
	fmt.Println(cheeses)
	//切片中删除元素
	cheeses = append(cheeses[:1], cheeses[2:]...)
	fmt.Println(cheeses)
	//复制切片中的元素
	var twoCheeses = make([] int, 2)
	copy(twoCheeses, cheeses)
	fmt.Println(twoCheeses)

	//内置函数make创建了一个映射，其键的类型为字符串，而值的类型为整数。
	var aboutMap = make(map[string] string)
	aboutMap["one"] = "one"
	aboutMap["two"] = "two"
	fmt.Println(aboutMap)

	delete(aboutMap, "one")
	fmt.Println(aboutMap)
}

//结构体和指针
//关键字type指定一种新类型
type Movie struct {
	Money int
	Name string
}

//嵌套结构体
type Two struct {
	one int
	two Movie
}

func useStruct()  {
	nowMovie := Movie{
		Money: 12,
		Name: "jojo",
	}

	fmt.Println(nowMovie)

	var twoMovie Movie
	twoMovie.Money = 22
	twoMovie.Name = "好"

	fmt.Println(twoMovie)

	three := new (Movie)
	three.Money = 12
	three.Name = "圣诞节"

	fmt.Println(three)

	four := Movie{
		Money: 32,
		Name: "iji",
	}

	fmt.Println(four)

	aboutTwo := Two{
		one: 12,
		two: Movie{
			Money: 32,
			Name: "iji",
		},
	}
	fmt.Println(aboutTwo)
}

//复制结构体
type ForCopy struct {
	one int
}

func aboutCopy()  {
	a := ForCopy{
		one: 12,
	}

	//复制值
	b := a
	b.one = 23
	fmt.Println(a, b)

	c := ForCopy{
		one: 11,
	}

	//复制地址
	d := &c

	d.one = 22

	fmt.Println(c, *d)
}


func main()  {
	aboutType()
	variable()
	pointer()

	i := 1
	showMemoryAddress(&i)

	fmt.Println(change)

	one, two := moreNum()
	fmt.Println(one, two)

	forMore(1, 2, 3)
	fmt.Println(knowName())

	fn := func() string {
		return "jojo"
	}

	fmt.Println(anotherFunction(fn))

	controlFlow()

	aboutDe()

	gather()

	useStruct()

	aboutCopy()
}
