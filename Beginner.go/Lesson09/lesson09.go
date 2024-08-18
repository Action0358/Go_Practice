package main

import "fmt"

func main() {
	x := 8
	y := 12
	z := 20

	// 複合代入演算子
	x += 10 // x = x + 10
	z += y  // z = z + y

	fmt.Println(x)
	fmt.Println(z)
}
