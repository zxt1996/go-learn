package readAndWrite

import (
	"fmt"
	"math/rand"
	"sync"
)

var count int
var wg sync.WaitGroup
var rw sync.RWMutex

func read(n int)  {
	rw.RLock()
	fmt.Printf("读取 %d\n", n)

	v := count

	fmt.Printf("读取 %d结束，值为: %d\n", n, v)
	wg.Done()
	rw.RUnlock()
}

func write(n int)  {
	rw.Lock()
	fmt.Printf("写 %d \n", n)
	v := rand.Intn(1000)

	count = v

	fmt.Printf("写 %d 结束，值为： %d \n", n, v)
	wg.Done()
	rw.Unlock()
}

func ReadAndWrite()  {
	wg.Add(10)
	for i:=0;i<5;i++ {
		go read(i)
	}

	for i:=0;i<5;i++ {
		go write(i)
	}

	wg.Wait()
}
