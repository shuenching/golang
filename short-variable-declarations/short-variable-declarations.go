package main

import "fmt"

//function外面不可使用短聲明
// test := 1

func main() {
	var i, j int = 1, 2
	//短聲明 := 可以用來替代var宣告，變數型態會依初始值自動給予
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
