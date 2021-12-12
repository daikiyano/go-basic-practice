package main

import "fmt"

type T struct {
	User
}

type User struct {
	Name string
	Age  int
	//X, Y int
}

//
//func (u User) SayName() {
//	fmt.Println(u.Name)
//}

//基本はポインタ型を使う
//
//func (u *User) SetName2(name string) {
//	u.Name = name
//}

func main() {
	//user1 := User{Name: "user1"}
	//user1.SayName()
	//
	//user1.SetName2("A")
	//user1.SayName()

	t := T{User: User{Name: "user1", Age: 10}}
	fmt.Println(t)
	fmt.Println(t.User)
	fmt.Println(t.User.Name)
	fmt.Println(t.Name)

}
