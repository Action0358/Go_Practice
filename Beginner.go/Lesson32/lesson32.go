package main

import "fmt"

func main() {
	var x interface{} = 3

	if x == nil { // x が nil かどうかをチェック
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt { // x が int 型かどうかをチェック
		fmt.Println("x is int", i)
	} else if s, isString := x.(string); isString { // x が string 型かどうかをチェック
		fmt.Println("x is string", s)
	} else {
		fmt.Println("I don't know") // x が他の型の場合
	}

}

// nil が初期値になる型
// var p *int       // ポインタ型。初期値は nil。
// var s []int      // スライス型。初期値は nil。
// var m map[string]int // マップ型。初期値は nil。
// var ch chan int  // チャネル型。初期値は nil。

// 出力結果
// fmt.Println(p)   // nil
// fmt.Println(s)   // nil
// fmt.Println(m)   // nil
// fmt.Println(ch)  // nil

// nil 以外の初期値を持つ型
// var i int        // 整数型。初期値は 0。
// var b bool       // ブール型。初期値は false。
// var str string   // 文字列型。初期値は空文字 ""。

// 出力結果
// fmt.Println(i)   // 0
// fmt.Println(b)   // false
// fmt.Println(str) // ""
