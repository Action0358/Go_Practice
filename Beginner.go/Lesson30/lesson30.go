package main

import "fmt"

func anything(a interface{}) { // インターフェース = 異なる型に共通の動作を定義するための契約
	fmt.Println(a) // 引数aとして定義したものを出力
}

func main() {
	anything("string") // 引数aにstringを渡す
	anything(1)        // 引数aに1を渡す

	var x interface{} = 3 // インターフェース
	i := x.(int)          // アサーション = インターフェース型から具体的な型を取り出す方法 (インターフェース型で変数定義された3をint型に復元する)
	fmt.Println(i + 2)    // int型に変換されたiのため、計算が可能
	// fmt.Println(x + 2) xではインターフェース型のため、型が異なり計算が不可能
}

// func main() {
// 	var x interface{} = 10 // インターフェース
// 	value, ok := x.(int)   // アサーション = インターフェース型の値が実際にはどの具体的な型を持っているかを確認したり、その具体的な型として扱ったりすること
// 	if ok {
// 		fmt.Printf("x is an int and its value is %d\n", value) // %dはフォーマット指定子で整数を10進数で表示するためのもの,\nは改行を意味する
// 	} else {
// 		fmt.Println("x is not an int")
// 	}
// }
