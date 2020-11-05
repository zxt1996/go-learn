package myErr

import (
	"errors"
	"fmt"
)

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

func MyErr()  {
	myErr()
}