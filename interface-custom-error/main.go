package main

import "fmt"

//エラーハンドリングの際にカスタムエラーを使用
//type error interface {
//	Error() string
//}

type MyError struct {
	Message string
	ErrCode int
}

// アドレスに格納されている値を返す
func (e *MyError) Error() string {
	return e.Message
}

//アドレスを参照して代入
func RaiseError() error {
	return &MyError{Message: "カスタムエラーが発生しました", ErrCode: 123}
}
func main() {
	err := RaiseError()
	fmt.Println(err.Error())
	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
	}
}
