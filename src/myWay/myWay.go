package myWay

import "fmt"

//方法
type Way struct {
	Name string
	Age int
}
//关键字func后面多了一个参数——接收者,方法接收者是一种类型
//只需编写方法实现一次，就可对结构体的任何实例进行调用。
func (m *Way) summary() string  {
	return m.Name + "好嘞"
}

func (m *Way) yourAge() int  {
	return m.Age - 1
}

func aboutWay()  {
	m := Way{
		Name: "sdk",
		Age: 12,
	}

	y := Way{
		Name: "神奇",
		Age: 22,
	}

	fmt.Println(m.summary())
	fmt.Println(m.yourAge())
	fmt.Println(y.summary())
}

//方法和指针
type Triangle struct {
	one int
}

func (m *Triangle) canChange() {
	m.one = 12
}

//没有在Triangle前面加上星号，这意味着接收者参数是值而不是指针
func (m Triangle) noChange()  {
	m.one = 13
	return
}

//如果需要修改原始结构体，就使用指针；如果需要操作结构体，但不想修改原始结构体，就使用值
func aboutChange()  {
	now := Triangle{
		one: 1,
	}

	fmt.Println(now)

	now.noChange()
	fmt.Println(now)

	now.canChange()
	fmt.Println(now)
}

func MyWay()  {
	aboutWay()

	aboutChange()
}
