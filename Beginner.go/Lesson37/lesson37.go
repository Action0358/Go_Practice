package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("example.txt") // ファイルを開く
	if err != nil {                     // エラーが発生したかどうかを確認
		log.Fatal(err) // エラーがあればエラーメッセージを出力し、プログラムを終了
	}
	defer file.Close() // ファイルを最後に必ず閉じる

	// file.Write([]byte("Hello")) example.txtにHelloを記述
}
