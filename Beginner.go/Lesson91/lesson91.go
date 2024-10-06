// sortパッケージは、スライスや配列をソート（並び替え）するための機能を提供

package main

import (
	"fmt"
	"sort"
)

func main() {
	// 整数のスライス
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}

	// スライスの内容を表示（ソート前）
	fmt.Println("Before sorting:", numbers)

	// 整数スライスをソート
	sort.Ints(numbers)

	// スライスの内容を表示（ソート後）
	fmt.Println("After sorting:", numbers)
}
