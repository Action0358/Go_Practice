package main

import (
	"fmt"
	"sort"
)

// Person構造体を定義
type Person struct {
	Name string
	Age  int
}

// Peopleという構造体のスライスを定義
type People []Person

// Lenメソッド（要素の数を返す）
func (p People) Len() int {
	return len(p)
}

// Lessメソッド（並び替えのルールを定義）
func (p People) Less(i, j int) bool { // 真偽値より並び替えの可否を判断
	return p[i].Age < p[j].Age
}

// Swapメソッド（要素を入れ替える）
func (p People) Swap(i, j int) { // 並び替えをSwapで行う
	p[i], p[j] = p[j], p[i]
}

func main() {
	// Personのスライスを作成
	people := People{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	// ソート前
	fmt.Println("Before sorting:", people)

	// 年齢でソート
	sort.Sort(people)

	// ソート後
	fmt.Println("After sorting:", people)
}
