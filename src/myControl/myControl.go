package myControl

import "fmt"

//控制流程
func controlFlow()  {
	can := true

	//if语句
	if can {
		fmt.Println("可以")
	}

	two := 12

	if two < 10 {
		fmt.Println("小于10")
	} else if two < 12 {
		fmt.Println("小于12")
	} else {
		fmt.Println("其他")
	}

	var all int
	for i := 0; i < 10; i++ {
		all += i
	}

	fmt.Println(all)

	//遍历数组
	numbers := []int {1, 2, 3}

	for index, value := range numbers{
		fmt.Println("numbers[", index, "]: ", value)
	}
}

//defer语句
//在函数返回前执行另一个函数
//通常用于执行清理操作或确保操作（如网络调用）完成后再执行另一个函数
func aboutDe()  {
	fmt.Println("start")

	defer fmt.Println("one")
	defer fmt.Println("two")

	fmt.Println("end")
}

func MyControl()  {
	controlFlow()

	aboutDe()
}