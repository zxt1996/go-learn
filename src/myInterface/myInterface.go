package myInterface

import (
	"fmt"
	"math"
)

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

func MyInterface()  {
	nowRect := rect{
		width: 1.2,
		height: 2.3,
		name: "zheling",
	}
	measure(nowRect)
}
