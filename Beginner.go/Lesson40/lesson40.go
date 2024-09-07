package main

import "fmt"

// グローバル変数
var count int

// init関数: プログラム開始前に初期化処理を行う
func init() {
	fmt.Println("Init function is called")
	count = 10 // countを初期化
}

// メイン関数: プログラムの開始点
func main() {
	fmt.Println("Main function is called")
	fmt.Println("Count value:", count) // initで設定されたcountの値を表示
}
