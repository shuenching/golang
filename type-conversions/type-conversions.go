package main

import (
	"fmt"
	"math"
)

//類型轉換Type(value)
//GO規定在類型轉換時一定要做轉型

func main() {
	var x, y int = 4, 3
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)
	fmt.Println(x, y, z)
}
