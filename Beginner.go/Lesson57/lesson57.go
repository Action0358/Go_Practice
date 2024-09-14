package main

import "fmt"

func main() {
	// 変数aを宣言して値を10に設定
	a := 10

	// 変数pはaのメモリアドレスを持つポインタ
	p := &a

	// aの値とそのメモリアドレスを表示
	fmt.Println("a:", a)   // aの値を表示
	fmt.Println("p:", p)   // aのメモリアドレス（ポインタ）を表示
	fmt.Println("*p:", *p) // pが指し示す値を表示（aの値）

	// ポインタpを使ってaの値を変更
	*p = 20

	// aの値が変更されたことを確認
	fmt.Println("a after p", p) // 変数の値が変更されてもメモリアドレス自体は変わらない
	fmt.Println("a after *p = 20:", a)
}
