package main

import (
	"fmt" // fmtパッケージをインポート
)

func main() {
	// 文字列の出力（改行あり）
	fmt.Println("こんにちは、Go言語！")

	// 変数の値を設定
	name := "Taro"
	age := 25
	height := 1.75

	// Printfはフォーマット指定子を使って変数を整形して出力
	// %s は文字列、%d は整数、%.2f は小数点以下2桁の浮動小数点数を意味する
	fmt.Printf("名前: %s\n", name)          // 名前を出力
	fmt.Printf("年齢: %d 歳\n", age)         // 年齢を出力
	fmt.Printf("身長: %.2f メートル\n", height) // 身長を小数点以下2桁で出力

	// Sprintfはフォーマットされた文字列を返し、変数に格納する
	intro := fmt.Sprintf("私は %s です。%d 歳です。", name, age)

	// 格納した文字列を出力
	fmt.Println(intro)
}
