package main

import "fmt"

func main() {

	//arrayとsliceの違い
	//sliceは要素数を指定する必要がない
	//var sl []int
	//fmt.Println(sl)
	//
	//var sl2 []int = []int{100, 200}
	//fmt.Println(sl2)
	//
	//sl3 := []string{"A", "B"}
	//fmt.Println(sl3)
	//
	//sl4 := make([]int, 5)
	//
	//fmt.Println(sl4)
	//
	//sl2[0] = 1000
	//fmt.Println(sl2)
	//
	//sl5 := []int{1, 2, 3, 4, 5}
	//fmt.Println(sl5)
	//fmt.Println(sl5[0])
	//fmt.Println(sl5[2:4])

	//sl := []int{100, 200}
	//fmt.Println(sl)
	//
	//sl = append(sl, 300)
	//fmt.Println(sl)
	//
	////make = sliceを生成
	//sl2 := make([]int, 5)
	//fmt.Println(sl2)
	//fmt.Println(len(sl2))
	//fmt.Println(cap(sl2))

	//スライス
	//copy

	//sl := []int{100, 200}
	//sl2 := sl
	//sl2[0] = 1000
	//fmt.Println(sl)

	//スライス　for
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)
	for i := range sl {
		fmt.Println(i, sl[i])
	}

}
