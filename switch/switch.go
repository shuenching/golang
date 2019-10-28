package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go runs on ")
	//runtime.GOOS是取得目前執行的作業系統
	//switch跟java和c一樣，不同的是break改成fallthrouth
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Println(os)
	}

	fmt.Println("===================")

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	//switch會由上至下一個一個檢查，如果任何條件成立，則使用default
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days")
	case today + 3:
		fmt.Println("In three days")
	case today + 4:
		fmt.Println("In four days")
	case today + 5:
		fmt.Println("In five days")
	default:
		fmt.Println("Too far away")
	}

	fmt.Println("===================")
	t := time.Now()
	//switch不帶條件，代表switch true，便於讓程式更清晰
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
