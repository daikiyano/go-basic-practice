package main

import (
	"fmt"
	"os"
)

//関数が終わった後にdeferを実行する
func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}

func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
}

func main() {
	TestDefer()

	//defer func() {
	//	fmt.Println("1")
	//	fmt.Println("2")
	//	fmt.Println("3")
	//}()

	RunDefer()

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file.Write([]byte("Hello"))
}
