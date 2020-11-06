package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

//一个安全的资源池，被管理的资源必须都实现io.Close接口
type Pool struct {
	m sync.Mutex  //互斥锁
	res chan io.Closer   //type Closer interface {Close() error}
	factory func()(io.Closer, error)
	closed bool
}

var ErrPoolClosed = errors.New("资源池已经关闭")

//创建一个资源池
//fn是创建新资源的函数；size是指定资源池的大小
func New(fn func()(io.Closer, error), size uint) (*Pool, error)  {
	if size <= 0 {
		return nil, errors.New("size的值太小")
	}

	return &Pool{
		factory: fn,
		res: make(chan io.Closer, size),  //使用size创建一个有缓冲的通道，来保存资源
	}, nil
}

//从资源池里获取一个资源
func (p *Pool) Acquire() (io.Closer, error)  {
	select {
	case r, ok := <- p.res:
		log.Println("Acquire:共享资源")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire:新生成资源")
		return p.factory()
	}
}

//关闭资源池，释放资源
func (p *Pool) Close()  {
	p.m.Lock()

	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	//关闭通道，不让写入
	close(p.res)  //关闭通道里的资源
	for r := range p.res {
		r.Close()
	}
}

//释放资源
func (p *Pool) Release(r io.Closer)  {
	//保证该操作和close方法的操作是安全的
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	case p.res <- r:
		log.Println("资源释放到池子里了")
	default:
		log.Println("资源池满了，释放资源")
		r.Close()
	}
}