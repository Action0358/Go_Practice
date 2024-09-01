package main

import "fmt"

func main() {
	var x interface{} = 3 // インターフェース型の変数 x に整数 3 を代入
	i := x.(int)          // インターフェース型 x を int 型にアサーションして i に代入
	fmt.Println(i + 2)    // i に 2 を加えた結果を出力（5）
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
