package main

import "fmt"

// キーが string 型、値が int 型のマップを宣言し、初期化
func main() {
	m := map[string]int{
		"Apple":  100, // キー "Apple" に対して値 100 を設定
		"Banana": 200, // キー "Banna" に対して値 200 を設定
	}

	// k = key, v = value
	for k, v := range m {
		fmt.Println(k, v)
	}
}
