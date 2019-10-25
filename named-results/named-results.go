package main

import "fmt"

//如果function寫了回傳值，但return沒有帶值的話，會直接返回當前有的變數
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
