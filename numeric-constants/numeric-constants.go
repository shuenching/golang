package main

import "fmt"

//常數會是一個高精度的值
const (
	Big   = 1 << 100
	Small = Big >> 99
)

//因常數未指定類型，所以會由上下文決定類型(下面例子依代入的參數型態和回傳值型態定義)
func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
