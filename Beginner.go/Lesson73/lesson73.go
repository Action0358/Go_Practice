// コマンドライン引数の取得

package main

// osパッケージは、OS（オペレーティングシステム）レベルの機能にアクセスするための標準ライブラリ
import (
	"fmt"
	"os"
)

func main() {
	// コマンドライン引数を取得
	args := os.Args

	// 引数の数を表示
	fmt.Println("引数の数:", len(args))

	// すべての引数を表示
	for i, arg := range args {
		fmt.Printf("引数%d: %s\n", i, arg)
	}
}

// 実行コマンド : go run lesson73.go first second third
