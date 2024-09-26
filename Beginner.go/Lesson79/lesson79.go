// logは、プログラムの動作状況やエラーを記録する
package main

import (
	"log"
	"os"
)

func main() {
	// 普通のログ出力
	log.Println("これは通常のログです")

	// エラーログ出力
	log.Printf("エラーが発生しました: %s", "ファイルが見つかりません")

	// エラーが発生した場合にログを出力してプログラムを終了
	_, err := os.Open("non-existent-file.txt")
	if err != nil { // nil は、「何もないこと」や「無効な状態」を表す値
		log.Fatal("致命的なエラー: ファイルを開けませんでした")
	}

	// Fatalではなく、Panicも使用可能
	// log.Panic("予期せぬエラーが発生しました")
}
