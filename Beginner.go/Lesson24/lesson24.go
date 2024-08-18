package main

import "fmt"

func cal(x, y int) (int, int) { // 戻り値のint型が2つ存在する場合、intを2つ分記述する
	return (x / y), (x * y)
}
func main() {
	result01, result02 := cal(6, 3) // 戻り値が2つ存在するため、結果も2つ記述する
	fmt.Println(result01, result02)
}

// 別の記述1
// func cal(x, y int) (int, int)
//		a := x / y
// 		b := x * y
// 		return a, b
// }

// 別の記述2
// func cal(x, y int) (a int, b int)
//		a := x / y
// 		b := x * y
// 		return
// }
