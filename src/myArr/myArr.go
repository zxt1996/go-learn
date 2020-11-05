package myArr

import "fmt"

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

func MyArr()  {
	gather()
}