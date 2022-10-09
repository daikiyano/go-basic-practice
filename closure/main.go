package main

import "fmt"

func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		// storeに上書きして保存する。次回呼び出すと上書きした値が取られる
		store = next
		return s
	}
}

func main() {
	f := Later()
	fmt.Println(f("Hello"))
	fmt.Println(f("My"))
	fmt.Println(f("name"))
	fmt.Println(f("is"))
}
