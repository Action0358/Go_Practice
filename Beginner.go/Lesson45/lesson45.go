package main

import "fmt"

func main() {
	// スライスを初期化。string型のスライスで、"A", "B", "C" という3つの要素を持つ
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)

	// forループを使用して、スライスの各要素にアクセス
	// rangeを使ってスライス sl のインデックスと要素を取得
	for i, v := range sl {
		// i はスライスのインデックス (位置)
		// v はそのインデックスにある値 (要素)
		fmt.Println(i, v) // 例: 0 A, 1 B, 2 C
	}
}
