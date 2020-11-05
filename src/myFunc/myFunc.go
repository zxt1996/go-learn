package myFunc

import "fmt"

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

func MyFunc()  {
	one, two := moreNum()
	fmt.Println(one, two)

	forMore(1, 2, 3)
	fmt.Println(knowName())

	fn := func() string {
		return "jojo"
	}

	fmt.Println(anotherFunction(fn))
}