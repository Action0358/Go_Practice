// ファイルやディレクトリの存在確認且つ作成と書き込み

package main

// osパッケージは、OS（オペレーティングシステム）レベルの機能にアクセスするための標準ライブラリ
import (
	"fmt"
	"os"
)

func main() {
	// チェックしたいファイルパス
	filename := "example.txt"

	// ファイルが存在するか確認
	// _はファイル情報（os.FileInfo 型）infoが該当するがファイル情報に興味ないため、無視
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		fmt.Println("ファイルが存在しません:", filename)
	} else {
		fmt.Println("ファイルが存在します:", filename)
	}

	// "example.txt" という名前のファイルを作成（すでに存在する場合は上書き）
	file, err := os.Create("example.txt")

	// エラーチェック
	if err != nil {
		fmt.Println("ファイル作成エラー:", err)
		return
	}

	// ファイルに書き込む
	// _はファイル情報
	_, err = file.WriteString("Hello, Go!\n")

	// エラーチェック
	if err != nil {
		fmt.Println("ファイル書き込みエラー:", err)
	}

	// ファイルを閉じる
	file.Close()

	fmt.Println("ファイル作成完了")
}
