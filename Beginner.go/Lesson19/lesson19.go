package main

import "fmt"

func sayHello() {
	fmt.Println("Hello Word!")
}

func main() {
	sayHello() // 定義した関数
	sayHello() // 複数回呼び出し可能
}
