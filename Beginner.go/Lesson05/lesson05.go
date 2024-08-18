package main

import "fmt"

func main() {
	a := [3]string{"sato", "suzuki", "takahashi"} // 変数名 := [要素数]データ型{配列データ1, 配列データ2, ...}
	a[1] = "tanaka"                               // インデックス1のデータを変更（suzuki→tanaka）

	b := [...]string{"sato", "suzuki", "takahashi"} //要素数の省略形式[...]

	c := [2][2]string{{"kagawa", "hashimoto"}, {"kadokawa", "miyamoto"}} // 多次元配列

	fmt.Println(a[0]) // インデックスは0からスタート
	fmt.Println(a[1])
	fmt.Println(a[2])

	fmt.Println(b[0])
	fmt.Println(b[1])
	fmt.Println(b[2])

	fmt.Println(c[0][0]) // 1グループ目の1番目
	fmt.Println(c[0][1]) // 1グループ目の2番目
	fmt.Println(c[1][0]) // 2グループ目の1番目
	fmt.Println(c[1][1]) // 2グループ目の2番目

}
