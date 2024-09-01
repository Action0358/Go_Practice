package main

import "fmt"

func anything(a interface{}) { // インターフェースとは異なる型に共通の動作を定義するための契約
	fmt.Println(a) // 引数aとして渡された値を出力
}

func main() {
	anything("string") // "string" を出力
	anything(1)        // 1 を出力
}
