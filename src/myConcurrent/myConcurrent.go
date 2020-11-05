package myConcurrent

import (
	"fmt"
	"time"
)

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

func MyConcurrent()  {
	//aboutGoroutine()

	nextGoroutine()

	twoChane()

	pingAndPong()

	aboutSelect()
}
