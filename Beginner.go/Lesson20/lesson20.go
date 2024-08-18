package main

import "fmt"

func sayHello(greeting string) {
	fmt.Println(greeting) // 引数の定義でstringを定義しているため、「""」は不要
}

func main() {
	sayHello("Good morning!")  // sayHello関数に"Good morning"を渡す
	sayHello("Good afternoon") // もし、sayHello関数に文字列型の引数を渡していないとエラーを出力する可能性がある
	sayHello("Good evening")
}
