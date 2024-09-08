package main

import "fmt"

func main() {
	// 整数型のチャネルを宣言
	var ch chan int
	// チャネルを初期化 (バッファなしのチャネル)
	ch = make(chan int)
	fmt.Println(cap(ch)) // cap 関数でチャネルのバッファ容量を確認 (バッファなしなので 0 が出力される)

	// バッファ付きのチャネルを作成（バッファサイズは5）
	ch2 := make(chan int, 5)
	fmt.Println(cap(ch2)) // バッファ容量 5 が出力される

	// チャネルに値を送信（バッファにデータを格納）
	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
	fmt.Println("len", len(ch2)) // len 関数で現在のチャネル内のデータ数を確認 (3 が出力される)

	// チャネルから値を受信（バッファからデータを取り出す）
	fmt.Println(<-ch2)           // 1 が出力される（先に送信したデータ）
	fmt.Println("len", len(ch2)) // チャネル内のデータ数が減り、2 になる

	fmt.Println(<-ch2)           // 2 が出力される（次のデータ）
	fmt.Println("len", len(ch2)) // チャネル内のデータ数が 1 に減る

	fmt.Println(<-ch2)           // 3 が出力される（最後のデータ）
	fmt.Println("len", len(ch2)) // チャネル内のデータ数が 0 になる
}
