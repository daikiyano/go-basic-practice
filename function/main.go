package main

import "fmt"

//func Plus(x int, y int) int {
//	return x + y
//}
//
//func Div(x, y int) (int, int) {
//	q := x / y
//	r := x % y
//	return q, r
//}
//
//func Double(price int) (result int) {
//	result = price * 2
//	return
//}

// 無名関数

//f := func(x, y int) int {
//	return x + y
//}
//
//i := f(1, 2)
//fmt.Println(i)

func ReturnFunc() func() {
	return func() {
		fmt.Println("I am a function")
	}
}

func main() {

	f := ReturnFunc()
	f()

	//i := Plus(1, 2)
	//fmt.Println(i)
	//
	//i2, _ := Div(9, 3)
	//
	//fmt.Println(i2)
	//
	//i4 := Double(1000)
	//fmt.Println(i4)
}
