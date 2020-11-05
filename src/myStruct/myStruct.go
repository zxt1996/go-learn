package myStruct

import "fmt"

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

func MyStruct()  {
	useStruct()

	aboutCopy()
}
