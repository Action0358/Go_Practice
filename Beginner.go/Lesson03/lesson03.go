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

	fmt.Println(string_a[0]) // 要素番号はbyte配列で出力される

	byteA := []byte{72, 69, 76, 76, 79}
	fmt.Println(byteA)

	fmt.Println(string(byteA)) // byte配列を文字型で出力するように変換するとHELLOになる

	byteB := []byte("HELLO")
	fmt.Println(byteB)

	fmt.Println(string(byteB))
}
