package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	//if 跟c或java用法一樣，但不使用()
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	//if 前面可以加一個簡單的語句，定義if作用域內可用的變數
	if v := math.Pow(x, n); v < lim {
		return v
	}
	//在if外使用v 會被告知v未定義
	//return v
	return lim
}

func pow2(x, n, lim float64) float64 {
	//使用if else作處理
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	//注意，這個function印出 27 >= 20\n 9 20
	//因為先執行了pow2(3, 2, 10) 再執行pow2(3, 3, 20) 最後才執行main的Println
	//所以先印了else的內容
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
}
