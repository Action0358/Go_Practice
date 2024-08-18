package main

import "fmt"

func cal(x int, y int) { // cal(x, y int) のように省略ができる
	fmt.Println(x / y)
}
func main() {
	cal(6, 3)
}
