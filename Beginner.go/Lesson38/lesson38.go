package main

import "fmt"

func main() {
	// deferで無名関数を定義し、main関数が終了する際に実行する処理を予約
	defer func() {
		// recover関数でパニック状態から回復を試みる
		if x := recover(); x != nil { // recover() がパニックメッセージをキャッチし、x に "runtime error" が代入される
			fmt.Println(x) // パニックの内容（今回の場合は"runtime error"）を出力
		}
	}()

	// panicで異常終了を発生させる
	panic("runtime error") // "runtime error" というメッセージでパニックを発生

	// この行には到達しない（panicによってプログラムはここで停止するため）
	fmt.Println("Start")
}
