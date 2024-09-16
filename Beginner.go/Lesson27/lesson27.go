package main

import "fmt"

type Student struct {
	name          string  // フィールド
	math, english float64 //　フィールド
}

// avg メソッド定義(変数名sを記述することで構造体のフィールドを使用することができる)、(s Student)をレシーバと言う
// avg() は 関数の一種（メソッド） であり、s Student というレシーバを通じて、Student 型のデータにアクセスできる
func (s Student) avg() {
	fmt.Println(s.name, (s.math + s.english/2))
}

func main() {
	a001 := Student{"sato", 80, 70}
	a001.avg()
}

// 別の記述1
// type Student struct {
// 		name  string
// }
//
// func (s Student) avg(math, english float64) {
// 	fmt.Println((s.math + s.english) / 2)
// }

// func main() {
// 	a001 := Student{"sato"}
// 	a001.avg(80, 70)
// }

// 別の記述2
// type Student struct {
// 		name  string
// }
//
// func (s Student) avg(math, english float64) float64 {
// 		return (s.math + s.english) / 2
// }

// func main() {
// 	a001 := Student{"sato"}
// 	fmt.Println(a001.avg(80, 70))
// }

// 別の記述3
// type Student struct {
// 		name  string
// }
//
// func (s Student) avg(math, english float64) (avgResult float64) {
// 		avgResult = (s.math + s.english) / 2
// 		return
// }

// func main() {
// 	a001 := Student{"sato"}
// 	fmt.Println(a001.avg(80, 70))
// }
