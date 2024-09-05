package main

import "fmt"

// ラベル付きforループの例

func main() {
Loop:
	for {
		for {
			for {
				fmt.Println("START") // 内側の最も深いループで "START" を出力
				break Loop           // "START" の後、ループ全体（Loopラベル付きのループ）を終了
			}
			fmt.Println("処理をしない") // ここには到達しない
		}
		fmt.Println("処理をしない") // ここにも到達しない
	}
	fmt.Println("END") // ここが表示される
}
