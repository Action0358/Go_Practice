package main

import "fmt"

func anything(a interface{}) { // インターフェースとは異なる型に共通の動作を定義するための契約
	switch v := a.(type) { // switchを用いて分岐処理を行うことができる
	case string:
		fmt.Println(v + "!?")
	case int:
		fmt.Println(v + 10000)
	}
}

func main() {
	anything("string") // "string!?" を出力
	anything(1)        // 10001 を出力
}
