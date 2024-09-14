// ポインタあり（参照渡し）
package main

import "fmt"

// 大きなデータの構造体
type LargeData struct {
	numbers [1000000]int
}

func processData(data *LargeData) {
	data.numbers[0] = 999 // ポインタを通じて直接データを変更
}

func main() {
	data := LargeData{}
	processData(&data)           // ポインタ（アドレス）を渡す
	fmt.Println(data.numbers[0]) // 999に変更されている
}

// 参照渡し:
// 関数にデータを渡す際に、データのメモリアドレスを渡す
// データのコピーを作成しないため、メモリの消費を抑えられる
// 関数内での変更が呼び出し元のデータに直接影響を与える（データの変更が関数外にも反映される）
