package main

import "fmt"

//function 可以接受多個參數或沒有參數

//兩個參數
func add(x int, y int) int {
	return x + y
}

//沒有參數
func noParams() string {
	return "no params"
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(noParams())
}
