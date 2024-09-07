package main

import "fmt"

func main() {
	sl := []string{"A", "B", "C"} // スライスを初期化。string型のスライスで、"A", "B", "C" という3つの要素を持つ
	fmt.Println(sl)

	// 通常のforループを使用して、スライスの各要素にアクセス
	// `i` はインデックスを表す変数。`len(sl)` でスライスの長さを取得して、ループの回数を決める
	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i]) // インデックス i に対応するスライスの要素 sl[i] を表示
	}
}
