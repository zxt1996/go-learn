package myArr

import "fmt"

//数组、切片和映射
func gather()  {
	var arr [2] int

	arr[0] = 1
	arr[1] = 2

	fmt.Println(arr)

	//内置函数make创建一个切片，其中第一个参数为数据类型，而第二个参数为长度
	//切片的底层是基于数组实现的
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

func newArr()  {
	slice := []int{1, 2, 3, 4, 5}

	copySlice := slice[1:3]

	copySlice = append(copySlice, 10)

	//因为newSlice有可用的容量，不会创建新的切片来满足追加，所以直接在newSlice后追加了一个元素10，
	//因为newSlice和slice切片共用一个底层数组，所以切片slice的对应的元素值也被改变了
	//所以一般我们在创建新切片的时候，最好要让新切片的长度和容量一样，这样我们在追加操作的时候就会生成新的底层数组，和原有数组分离，
	fmt.Println(slice)
	fmt.Println(copySlice)
}

func MyArr()  {
	//gather()
	newArr()
}