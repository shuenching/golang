package main

import (
	"fmt"
	"math/cmplx"
)

//go 基本型態
//bool 布林
//string 字串
//int int8 int16 int32 int64 整數
//uint uint8 uint16 uint64 uintptr 無號數
//byte uint8的別名，代表一個位元組
//rune int32的別名，代表一個unicode code
//float32 float64 浮點數
//complex64 complex128

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
