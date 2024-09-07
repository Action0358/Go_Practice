package main

import (
	"fmt"
	"time"
)

// sub関数: 無限ループで「Sub Loop」を表示する
// この関数は後で並行処理（goroutine）で実行されます
func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond) // 100ミリ秒（0.1秒）待機
	}
}

func main() {
	go sub() // sub関数をgoroutineとして並行処理で実行

	// 無限ループで「Main Loop」を表示する
	for {
		fmt.Println("Main Loop")
		time.Sleep(200 * time.Millisecond) // 200ミリ秒（0.2秒）待機
	}
}
