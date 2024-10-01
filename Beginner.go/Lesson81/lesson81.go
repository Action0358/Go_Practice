// strings パッケージは、文字列の操作（検索、置換、分割、結合など）を扱う便利な関数が集まった標準パッケージ

package main

import (
	"fmt"
	"strings"
)

func main() {
	// 文字列の中に特定の部分文字列が含まれているかチェック
	sentence := "Hello, Go!"
	containsGo := strings.Contains(sentence, "Go") // Contains : 結果を真偽値で返す
	fmt.Println("文字列に 'Go' が含まれている:", containsGo)

	// 文字列の前後の空白を取り除く
	spacedStr := "  Go Programming  "
	trimmedStr := strings.TrimSpace(spacedStr) // TrimSpace : 前後の空白を除去
	fmt.Println("前後の空白が取り除かれた文字列:", trimmedStr)

	// 文字列の置換
	replacedStr := strings.Replace(sentence, "Go", "World", 1) // Replace : 指定部分を置換
	fmt.Println("置換後の文字列:", replacedStr)

	// 文字列の分割
	str := "a,b,c"
	splitStr := strings.Split(str, ",") // Split : 区切り文字で分割
	fmt.Println("文字列の分割:", splitStr)

	// 文字列の結合
	joinedStr := strings.Join(splitStr, "-") // Join : 区切り文字で結合
	fmt.Println("結合された文字列:", joinedStr)

	// 大文字・小文字の変換
	upperStr := strings.ToUpper(sentence) // ToUpper : 文字列をすべて大文字
	lowerStr := strings.ToLower(sentence) // ToLower : 文字列をすべて小文字
	fmt.Println("大文字に変換:", upperStr)
	fmt.Println("小文字に変換:", lowerStr)
}
