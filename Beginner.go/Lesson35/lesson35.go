package main

import "fmt"

func main() {
Loop:
	for i := 0; i < 3; i++ { // 外側のループ。i は 0 から 2 まで繰り返す
		for j := 1; j < 3; j++ { // 内側のループ。j は 1 から 2 まで繰り返す
			if j > 1 { // j が 1 より大きい場合
				continue Loop // 外側のループ (Loop) の次の繰り返しにスキップする
			}
			// j が 1 のときのみ実行される
			fmt.Println(i, j, i*j) // i と j の値、および i と j の積を表示
		}
		// この行には到達しない (j > 1 の場合)
		fmt.Println("処理をしない") // 内側のループのすべての繰り返しが終わった後に表示される
	}
}
