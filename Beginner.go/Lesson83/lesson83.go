package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// test.txtを作成
	file, err := os.Create("test.txt") // test.txtを作成し、fileにその参照を格納
	if err != nil {                    // エラーが発生した場合
		log.Fatal(err) // エラーログを出力し、プログラムを終了
	}
	defer file.Close() // 関数が終了する際にファイルを閉じる

	// ファイルに文字列を書き込む
	_, err = file.WriteString("Hello, Go!") // ファイルに文字列を書き込む
	if err != nil {                         // エラーが発生した場合
		log.Fatal(err) // エラーログを出力し、プログラムを終了
	}

	fmt.Println("ファイルが正常に作成されました") // 成功メッセージを表示
}
