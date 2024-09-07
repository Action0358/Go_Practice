package main

import "fmt"

func main() {
	// キーが string 型、値が int 型のマップを宣言し、初期化
	m := map[string]int{
		"Apple": 100, // キー "Apple" に対して値 100 を設定
		"Banna": 200, // キー "Banna" に対して値 200 を設定
	}

	// マップ m のすべてのキーを繰り返し処理
	for k := range m {
		fmt.Println(k)
	}
}
