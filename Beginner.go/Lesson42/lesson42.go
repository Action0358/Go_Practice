package main

import "fmt"

// append

func main() {
	sl := []int{100, 200} // []int{100, 200} という2つの要素を持つ int 型のスライスを作成
	fmt.Println(sl)

	sl = append(sl, 300) // append 関数でスライスに新しい要素300を追加
	fmt.Println(sl)
}
