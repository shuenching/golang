package main

import "fmt"

//在function中如果是同類型的連續參數，可以只在最後一個宣告型態
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
