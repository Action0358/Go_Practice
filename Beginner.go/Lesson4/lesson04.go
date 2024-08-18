package main

import (
	"fmt"
	"reflect" // プログラム実行時に型情報や値の詳細を動的に調べたり操作したりするための機能
)

func main() {
	a := 10
	b := 1

	var num_bool bool = a > b

	fmt.Println(num_bool)
	fmt.Println(reflect.TypeOf(num_bool))
}
