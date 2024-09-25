// Flag : プログラムの動作を制御するためにコマンドラインから渡すオプション引数

package main

import (
	"flag" // flagパッケージをインポート
	"fmt"
)

func main() {
	// フラグの定義
	// -name フラグ: 名前を指定する。デフォルト値は "Go言語ユーザー"
	name := flag.String("name", "Go言語ユーザー", "あなたの名前")

	// -age フラグ: 年齢を指定する。デフォルト値は 20
	age := flag.Int("age", 20, "あなたの年齢")

	// -verbose フラグ: 詳細出力を有効にするかどうかを指定する。デフォルト値は false
	verbose := flag.Bool("verbose", false, "詳細出力を有効にする")

	// フラグをパース（解析）する
	// コマンドライン引数を元に定義したフラグの値を設定する（ macOSやLinuxでは「ターミナル」）
	// name = 任意な引数, age = 任意な引数, verbose = -verbose（引数）
	flag.Parse()

	// フラグの値を使用して挨拶メッセージを出力
	fmt.Printf("こんにちは、%sさん！\n", *name)

	// 年齢を出力
	fmt.Printf("年齢: %d\n", *age)

	// -verbose フラグが true の場合、詳細出力を行う
	if *verbose {
		fmt.Println("詳細出力モードが有効です。")
	}
}

// 実行コマンド : go run lesson77.go -name="Taro" -age=25 -verbose
// コマンドライン引数の前につける「-」は、フラグやオプションを指定するための記号
