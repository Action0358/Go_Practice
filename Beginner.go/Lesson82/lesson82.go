// bufio パッケージは、入力と出力をバッファリングするために使用される

package main

import (
	"bufio"
	"fmt"
	"os"
)

// bufio パッケージを使ったファイルの読み取り

func main() {
	// ファイルを開くための関数
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("ファイルを開けませんでした:", err)
		return
	}
	defer file.Close() // deferは関数の終了直前に実行される処理を予約する機能

	// バッファを使ってファイルを1行ずつ読み取る
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("読み取った行:", line) // 読み取った行を表示
	}

	// エラーチェック
	if err := scanner.Err(); err != nil { // scanner.Err() はファイルの読み取り中に発生したエラーを返すメソッド
		fmt.Println("読み取り中にエラーが発生しました:", err)
	}
}
