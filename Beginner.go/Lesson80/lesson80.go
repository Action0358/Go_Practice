package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 整数から文字列への変換
	i := 123
	s := strconv.Itoa(i) // int を string に変換
	fmt.Println("int to string:", s)

	// 文字列から整数への変換
	str := "456"
	num, err := strconv.Atoi(str) // string を int に変換
	if err != nil {
		fmt.Println("変換に失敗しました:", err)
	} else {
		fmt.Println("string to int:", num)
	}

	// 文字列から浮動小数点数への変換
	fStr := "3.14159"
	f, err := strconv.ParseFloat(fStr, 64) // string を float64 に変換
	if err != nil {
		fmt.Println("浮動小数点数の変換に失敗しました:", err)
	} else {
		fmt.Println("string to float:", f)
	}

	// 整数を16進数文字列に変換
	hexStr := strconv.FormatInt(255, 16) // int64 を 16進数表現の string に変換
	fmt.Println("int to hex string:", hexStr)
}
