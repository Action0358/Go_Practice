package main

import (
	"fmt"
	"sort"
)

func main() {
	// 文字列のスライス
	strings := []string{"banana", "apple", "cherry", "date"}

	// ソート前
	fmt.Println("Before sorting:", strings)

	// 文字列スライスをソート
	sort.Strings(strings)

	// ソート後
	fmt.Println("After sorting:", strings)
}
