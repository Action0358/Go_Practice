package main

import (
	"fmt"
	"regexp"
)

// 文字列の分割

func main() {
	// 正規表現パターン: スペースで分割
	re := regexp.MustCompile(`\s+`)

	// 対象文字列
	text := "Go is an open-source programming language"

	// 正規表現を使って分割
	words := re.Split(text, -1) // 分割回数の指定

	fmt.Println("分割結果:", words)
}

/*
  パターン定義: \s+
  \s+: 1つ以上の空白文字（スペース、タブ、改行など）にマッチ

  Split 関数:
  対象文字列を指定した正規表現パターンに基づいて分割します。ここでは、空白文字を区切りとして使用

  n > 0: 指定された回数分だけ文字列を分割
  n == 0: 何も返さない（空のスライスを返す）
  n < 0: 分割回数に制限を設けず、全てのマッチ箇所で分割
*/
