package main

import "fmt"

func main() {
	sum := 0
	//GO只有一種迴圈，跟c或java的差別是不使用括號()
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//Go的迴圈也可以不使用前置(起始值)、後置(遞增量)語話
	//for ;條件;{}  像這樣的寫法，但lint移除了前後分號
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	//無窮迴圈，不加任何條件的for即為無窮迴圈
	//for {}
}
