package main

import (
	"fmt"
	"regexp"
)

// パターンを使った文字列の置換

func main() {
	// 正規表現パターン: 数字を全て置換
	re := regexp.MustCompile(`[0-9]+`)

	// 対象文字列
	text := "ID12345 and IDabc"

	// 数字を「X」で置換
	result := re.ReplaceAllString(text, "X")

	fmt.Println("置換後のテキスト:", result)
}

// パターン定義: [0-9]+
// [0-9]+: 1文字以上の数字にマッチ。

// ReplaceAllString 関数:
// 対象の文字列内でパターンに一致する部分を「X」に置換。
