package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5} // []int{1, 2, 3, 4, 5} で要素を持つスライスを作成
	sl2 := make([]int, 5, 10)  // この時点では全ての要素が 0 になっている
	fmt.Println(sl2)

	n := copy(sl2, sl)  // copy関数で sl の内容を sl2 にコピー
	fmt.Println(n, sl2) // コピーされた要素の数 n と、コピー後の sl2 の内容を表示
}
