package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// ゴルーチンでch1にデータを送信
	go func() {
		for {
			ch1 <- "data from ch1"
			time.Sleep(1 * time.Second)
		}
	}()

	// ゴルーチンでch2にデータを送信
	go func() {
		for {
			ch2 <- "data from ch2"
			time.Sleep(1 * time.Second)
		}
	}()

	// 5回 select 文を実行してランダム性を確認
	for i := 0; i < 5; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		}
	}
}
