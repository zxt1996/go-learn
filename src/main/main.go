package main

import (
	"errors"
	"fmt"
	"math"
	"time"
	"myType"
	"myPointer"
	"myFunc"
	"myControl"
	"myArr"
	"myStruct"
)





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

//接口
type geometry interface {
	area() string
	per() float64
}

type rect struct {
	width, height float64
	name string
}

func (r rect) area() string  {
	return "here is " + r.name
}

func (r rect) per() float64  {
	return math.Pi * r.height * r.width
}

func measure(g geometry)  {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.per())
}

//错误处理
func f1(arg int) (int, error)  {
	if arg == 42 {
		//使用errors.New 可返回一个错误信息
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

type argError struct {
	arg int
	prob string
}

//fmt包在处理error时会调用Error方法
func (e *argError) Error() string  {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error)  {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func myErr()  {
	for _, i := range [] int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}

		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
}

//并发和Goroutine
func slowFunc()  {
	time.Sleep(time.Second * 2)
	fmt.Println("sleeper finished")
}

func aboutGoroutine()  {
	//只需在要让Goroutine执行的函数或方法前加上关键字go即可
	go slowFunc()
	fmt.Println("end")
	time.Sleep(time.Second * 3)
}

//通道
func slowTwo(c chan string)  {
	time.Sleep(time.Second)
	c <- "finished"
}


func nextGoroutine()  {
	c := make(chan string)

	go slowTwo(c)

	fmt.Println("next")
	msg := <- c
	fmt.Println("get mas")
	fmt.Println(msg)
}

//缓冲通道
func twoChane()  {
	//第二个参数表示表示缓冲区长度
	two := make(chan string, 2)

	two <- "one"
	two <- "two"

	fmt.Println(<-two)
	fmt.Println("一有数据传送过来立马接收")
	fmt.Println(<-two)
}

//<-位于关键字chan左边时，表示通道在函数内是只读的；
//<-位于关键字chan右边时，表示通道在函数内是只写的；
//没有指定<-时，表示通道是可读写的。
func ping(pings chan <- string, msg string)  {
	pings <- msg
}

func pong(pings <- chan string, pongs chan <- string)  {
	msg := <- pings
	pongs <- msg
}

func pingAndPong()  {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<- pongs)
	return
}

func aboutSelect()  {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "小于10"
	}()

	go func() {
		time.Sleep(time.Second)
		c2 <- "大于"
	}()

	select {
	case mes1 := <- c1:
		fmt.Println(mes1)
	case mes2 := <- c2:
		fmt.Println(mes2)
		//使用超时时间。这让select语句在指定时间后不再阻塞，以便接着往下执行
	case <- time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

func main()  {
	myType.MyType()

	myPointer.MyPointer()

	myFunc.MyFunc()

	myControl.MyControl()

	myArr.MyArr()

	myStruct.MyStruct()

	aboutWay()

	aboutChange()

	nowRect := rect{
		width: 1.2,
		height: 2.3,
		name: "zheling",
	}
	measure(nowRect)

	myErr()

	//aboutGoroutine()

	nextGoroutine()

	twoChane()

	pingAndPong()

	aboutSelect()

}
