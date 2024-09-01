package main

import "fmt"

func main() {
	x := "apple"
	switch x { // switch = 複数の条件に基づいて分岐処理を行う制御構文
	case "apple", "banana", "peach":
		fmt.Println("x is a fruit")
	case "carrot", "celery", "beet":
		fmt.Println("x is a vegetable")
	default:
		fmt.Println("x is not a fruit or a vegetable")
	}

	n := 6
	switch {
	case n > 0 && n < 4:
		fmt.Println("0 < n < 4")
	case n > 3 && n < 7:
		fmt.Println("3 < n <7")
	}
}
