package main

import "fmt"

func main() {
	// マップの宣言と初期化
	// キーが string 型で、値が int 型のマップを作成
	var m = map[string]int{"A": 100, "B": 200}
	// マップ全体を出力
	fmt.Println(m)

	// マップからキー "A" の値を取得して出力
	fmt.Println(m["A"])
	// マップからキー "B" の値を取得して出力
	fmt.Println(m["B"])

	// 略式でのマップ宣言と初期化
	m2 := map[string]int{"A": 100, "B": 200}
	// m2 マップを出力
	fmt.Println(m2)

	// キーが int 型で、値が string 型のマップを宣言
	m3 := map[int]string{
		1: "A", // キー 1 に対して値 "A"
		2: "B", // キー 2 に対して値 "B"
	}
	// m3 マップを出力
	fmt.Println(m3)

	// キー 1 の値を取得して出力
	fmt.Println(m3[1])
	// キー 2 の値を取得して出力
	fmt.Println(m3[2])

	// make関数を使用して、空のマップ m4 を作成（キーが int、値が string）
	m4 := make(map[int]string)
	// 空のマップ m4 を出力
	fmt.Println(m4)

	// マップ m4 にキーと値を追加
	m4[1] = "A" // キー 1 に "A" を設定
	m4[2] = "B" // キー 2 に "B" を設定
	// 追加後のマップ m4 を出力
	fmt.Println(m4)

	// マップ m4 からキー 3 の値を取得し、存在するかをチェック
	_, ok := m4[3] // k = key, v = value
	// キー 3 が存在しない場合の処理
	if !ok {
		fmt.Println("error") // キー 3 が存在しない場合 "error" と出力
	}

	// キー 2 をマップから削除
	delete(m4, 2)
	fmt.Println(m4) // キー 2 が削除されたマップを出力

	// マップの要素数を取得して出力
	fmt.Println(len(m4)) // マップ m4 の要素数を出力
}
