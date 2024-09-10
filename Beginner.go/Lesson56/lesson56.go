package main

import (
	"fmt"
	"time"
)

func main() {
	// 文字列型を送受信するチャネルを作成
	ch := make(chan string)

	// 別のゴルーチンで3秒後にデータを送信
	go func() {
		time.Sleep(3 * time.Second) // 3秒待機
		ch <- "data from channel"   // チャネルにデータを送信
	}()

	// select文を使ってチャネルからのデータ受信を待機
	select {
	case msg := <-ch:
		// チャネルからデータが受信された場合に実行
		fmt.Println(msg)
	case <-time.After(2 * time.Second):
		// 2秒経過してもデータが受信されなかった場合にタイムアウトとして実行
		fmt.Println("timeout")
	}
}
