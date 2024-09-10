package main

import (
	"fmt"
	"time"
)

// receive 関数は、チャネル ch からデータを受信して出力する
func receive(ch chan int) { // 整数型データを送受信するチャネル ch を引数として定義する
	for {
		// チャネルからデータを受信し、変数 i に格納
		i := <-ch
		// 受信したデータを出力
		fmt.Println(i)
	}
}

// main 関数は、チャネル ch にデータを送信する
func main() {
	// 整数型のチャネルを作成
	ch := make(chan int)

	// 別のゴルーチンで receive 関数を実行し、並列処理を開始
	go receive(ch) // receive 関数がバックグラウンドで並行して動作

	i := 0
	for i < 100 {
		// チャネルに i の値を送信
		ch <- i
		// 50 マイクロ秒待機
		time.Sleep(50 * time.Microsecond)
		i++
	}
}

// 関数の引数名はその関数のスコープ内だけで有効。つまり、各関数で使われる引数名は独立している
// 可読性向上のため、データ送受信の引数chを統一している
