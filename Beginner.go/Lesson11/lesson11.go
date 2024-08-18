package main

import "fmt"

func main() {
	age := 22

	if age >= 20 { // 条件分岐
		fmt.Println("adult")
	} else if age == 0 {
		fmt.Println("baby")
	} else {
		fmt.Println("child")
	}
}
