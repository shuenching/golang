package main

import "fmt"

//Vertex defined X-axis and Y-axis 定義一個類
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
