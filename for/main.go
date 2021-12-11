package main

import "fmt"

func main() {
	//point := 0
	//for point < 10 {
	//	fmt.Println(point)
	//	point++
	//}

	//for i := 0; i < 10; i++ {
	//	if i == 3 {
	//		continue
	//	}
	//	fmt.Println(i)
	//}

	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
