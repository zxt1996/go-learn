package myConcurrent

import (
	"fmt"
	"sync"
	"time"
)

//并发和Goroutine
//go语言中并发指的是让某个函数独立于其他函数运行的能力，一个goroutine就是一个独立的工作单元
//当我们创建一个goroutine的后，会先存放在全局运行队列中，等待Go运行时的调度器进行调度，把他们分配给其中的一个逻辑处理器，并放到这个逻辑处理器对应的本地运行队列中，最终等着被逻辑处理器执行即可。
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

func aboutGo()  {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i:=1;i<12;i++ {
			fmt.Println("A: ", i)
		}
	}()

	go func() {
		defer wg.Done()
		for i:=1;i<12;i++ {
			fmt.Println("B: ", i)
		}
	}()

	wg.Wait()
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

//无缓冲通道
//发送goroutine和接收gouroutine必须是同步的，同时准备后，如果没有同时准备好的话，先执行的操作就会阻塞等待，直到另一个相对应的操作准备好为止。
//这种无缓冲的通道我们也称之为同步通道
func noStore()  {
	ch := make(chan int)

	go func() {
		var sum int = 0

		for i :=0 ; i < 10;i++ {
			sum += i
		}

		ch <- sum
	}()

	fmt.Println(<- ch)
}


//缓冲通道
//对于有缓冲的通道，向其发送操作就是向队列的尾部插入元素，接收操作则是从队列的头部删除元素，并返回这个刚刚删除的元素
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

	//aboutGo()

	//nextGoroutine()

	noStore()

	//twoChane()
	//
	//pingAndPong()
	//
	//aboutSelect()
}
