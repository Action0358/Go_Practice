// ポインタなし（値渡し）
package main

import "fmt"

// 大きなデータの構造体
type LargeData struct {
	numbers [1000000]int
}

func processData(data LargeData) {
	data.numbers[0] = 999 // この変更は元のデータには反映されない
}

func main() {
	data := LargeData{}
	processData(data)            // コピーが作られる
	fmt.Println(data.numbers[0]) // 0のまま
}

// 値渡し:
// 関数にデータを渡す際に、データのコピーが作成される
// 大きなデータ（例: 大きな構造体や配列）を渡すと、多くのメモリを消費する
// コピーが作成されるため、関数内での変更が呼び出し元のデータに影響を与えない
