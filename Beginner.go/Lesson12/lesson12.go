package main

import "fmt"

func main() {
	x := 10
	y := 12

	if age := x + y; age >= 20 { // 簡易文＝if文の中に変数を宣言、代入、四則演算の簡易的な文を記述できる
		fmt.Println("adult", age)
	} else if age == 0 {
		fmt.Println("baby")
	} else {
		fmt.Println("child")
	}
}
