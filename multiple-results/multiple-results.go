package main

import "fmt"

//function可以返回多個參數，在function後用()指定回傳類型
func swap(x, y string) (string, string) {
	return y, x
}

//接收回傳值時，要依function的回傳數是定義
func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
