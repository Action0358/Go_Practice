package main

import "fmt"

// 可変長引数を受け取る関数
// 引数 nums はスライスとして扱われる
func Sum(nums ...int) int { // 可変長引数の構文 : func 名前(引数名 ...型)
	total := 0

	// k = key, v = value
	for _, v := range nums { // キーが存在しない場合の処理
		total += v
	}
	return total
}

func main() {
	fmt.Println(Sum(1, 2, 3))
	fmt.Println(Sum(10, 20, 30, 40))
}
