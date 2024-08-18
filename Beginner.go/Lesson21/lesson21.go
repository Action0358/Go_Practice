package main

import "fmt"

func cal(x int) { // 関数定義にcalculate関数を採用し、int型のxを引数に渡す
	fmt.Println(x * 3) // 6 * 3
}
func main() {
	cal(6)
}
