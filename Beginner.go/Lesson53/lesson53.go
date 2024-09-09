package main

import "fmt"

// データをチャネルに送信する関数
func sendData(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i // チャネルにデータを送信
	}
	close(ch) // 送信が終わったらチャネルを閉じる
}

func main() {
	ch := make(chan int)

	go sendData(ch) // ゴルーチンでsendData関数を実行

	// チャネルからデータを受信し続ける
	for value := range ch {
		fmt.Println(value) // 受信したデータを出力
	}

	fmt.Println("チャネルが閉じられました")
}
