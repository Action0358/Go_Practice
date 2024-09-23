package main

import "fmt"

func main() {
	x := 10 // x はmain関数内でのみ有効
	fmt.Println("Outer x:", x)

	if true {
		x := 20 // これは新しい変数 x（スコープは if ブロック内）
		fmt.Println("Inner x:", x)
	}

	fmt.Println("Outer x after if block:", x)

	for i := 0; i < 3; i++ {
		y := i * 2 // y はforループ内でのみ有効
		fmt.Println("y inside loop:", y)
	}

	// fmt.Println(y) // これはエラーになる。yはforループの外では無効。
}

// パッケージスコープ
// ファイル全体で有効
// 大文字で始まる識別子は他のパッケージからもアクセス可能

// ブロックスコープ
// 関数、if文、forループなどのブロック内でのみ有効
// 内側で同じ名前の変数を宣言すると、内側の変数が優先される（シャドーイング）
