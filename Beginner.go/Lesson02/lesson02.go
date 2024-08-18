package main

import (
	"fmt"
	"reflect" // プログラム実行時に型情報や値の詳細を動的に調べたり操作したりするための機能
)

func main() {
	num1 := 123
	var num2 int = 1234567890
	num3 := 1.23
	var num4 float64 = 1.23456789

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
	fmt.Println(reflect.TypeOf(num1)) //reflect.TypeOf関数より型の情報を表す
	fmt.Println(reflect.TypeOf(num2))
	fmt.Println(reflect.TypeOf(num3))
	fmt.Println(reflect.TypeOf(num4))
}
