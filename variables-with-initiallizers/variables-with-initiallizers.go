package main

import "fmt"

//變數初始化可以使用下面的表達方式，在型態後加=給值的順序跟變數宣告順序相同
var i, j int = 1, 2
var c, python, java = true, false, "no!"

func main() {
	fmt.Println(i, j, c, python, java)
}
