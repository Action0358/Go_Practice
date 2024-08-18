package main

import "fmt"

type Student struct { // 構造体を定義、Studentは構造体名
	name    string  // 6~7行目はフィールド定義(生徒の名前を格納するnameと数学と英語の点数を格納するmathとenglish)
	math    float64 // math, english float64のように省略できる
	english float64
}

func main() {
	var s Student // 構造体の初期化(変数定義、変数名s、構造体名Student)

	s.name = "sato"
	s.math = 80
	s.english = 70

	fmt.Println(s)
}

//別の記述
// func main() {
//		s := Student{"sato", 80, 70}
// 		fmt.Println(s)
// }

// 一部のみフィールド指定
// func main() {
// 		s := Student{name: "sato"}
// }
