package main

import "fmt"

func main() {
	hello := func(greeting string) { // funcの後に関数名がないことを無名関数
		fmt.Println(greeting)
	}
	hello("Good morning")
}

// func main() {
//		func(greeting string) {
// 			fmt.Println(greeting)
// 		}("Good morning")
// }
