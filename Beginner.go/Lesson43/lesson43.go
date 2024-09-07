package main

import "fmt"

// make関数を使って、長さ5のint型のスライスを作成

func main() {
	sl := make([]int, 5) // スライスの要素はデフォルトで全て0に初期化される
	fmt.Println(sl)

	fmt.Println(len(sl)) // len関数を使ってスライスの長さを表示

	fmt.Println(cap(sl)) // cap関数を使ってスライスの容量を表示

	sl2 := make([]int, 5, 10) // make関数を使って、長さ5、容量10のint型スライスを作成
	fmt.Println(sl2)

	fmt.Println(len(sl2)) // スライスの長さを表示

	fmt.Println(cap(sl2)) // スライスの容量を表示

	sl2 = append(sl2, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(sl2))

	fmt.Println(cap(sl2)) // 容量を超えると倍増する
}
