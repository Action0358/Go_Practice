package main

import "fmt"

func main() {
	x := 10
	y := 2

	fmt.Println(x >= 5 && x <= 10) // 論理演算子 and
	fmt.Println(y >= 5 && y <= 10)

	fmt.Println(x == 2 || y == 2) // 論理演算子 or
	fmt.Println(x == 1 || y == 1)
}
