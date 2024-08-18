package main

import "fmt"

func cal(x, y int) int { // 戻り値のデータ型を記述
	return (x / y)
}
func main() {
	result := cal(6, 3)
	fmt.Println(result)
}
