package main

import (
	"fmt"
	"reflect" // プログラム実行時に型情報や値の詳細を動的に調べたり操作したりするための機能
)

func main() {
	var string_a string = "Hello World"
	string_b := "Hello World!"

	fmt.Println(string_a)
	fmt.Println(string_b)
	fmt.Println(reflect.TypeOf(string_a))
	fmt.Println(reflect.TypeOf(string_b))
}
