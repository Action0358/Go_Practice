package main

import "fmt"

// deferは関数の終了直前に実行される処理を予約する機能

func TestDefer() {
	defer fmt.Println("END") // この行の実行が後回しになる
	fmt.Println("START")     // まずこの行が実行される
}

func main() {
	TestDefer()

	defer func() { // 複数の処理をまとめた
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
	}()
}
