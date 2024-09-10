package main

import "fmt"

func main() {
	// バッファサイズ3の整数型チャネルを作成
	ch := make(chan int, 3)

	// チャネルに3つの整数を送信
	ch <- 1
	ch <- 2
	ch <- 3

	// チャネルを閉じる（これ以上データが送信されないことを通知）
	close(ch)

	// rangeループを使用して、チャネルからデータを受信
	for i := range ch {
		// 受信したデータを出力
		fmt.Println(i)
	}
}
