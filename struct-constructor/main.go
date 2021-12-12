package main

import "fmt"

type User struct {
	Name string
	Age  int
}

//func NewUser(name string, age int) *User {
//	return &User{Name: name, Age: age}
//}
//
//func main() {
//	user1 := NewUser("user1", 23)
//	fmt.Println(user1)
//}

type Users []*User

func main() {
	user1 := User{Name: "user1", Age: 10}
	user2 := User{Name: "user2", Age: 20}
	user3 := User{Name: "user3", Age: 30}
	user4 := User{Name: "user4", Age: 40}

	users := Users{}
	users = append(users, &user1)
	users = append(users, &user2)
	users = append(users, &user3)
	users = append(users, &user4)
	fmt.Println(users)

	for _, u := range users {
		fmt.Println(*u)
	}

}
