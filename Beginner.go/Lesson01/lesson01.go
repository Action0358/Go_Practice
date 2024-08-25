package main // Goのプログラムが実行される起点となるような特別なパッケージ

import "fmt" // fmt パッケージに用意された Print 系関数

func main() { // main関数を宣言
	var num int      // var=変数宣言　num=変数名を記述　int=変数に代入するデータの種類
	num = 1          // 変数に引数を代入
	fmt.Println(num) // 変数に代入した値を表示

	var num1 = 2 // int型を省略した定義、Goが自動的にデータ型を識別してくれる
	fmt.Println(num1)

	num2 := 3 // 変数宣言と同時に代入した定義、":="はGoでよく使用される変数の宣言方法
	fmt.Println(num2)
}
