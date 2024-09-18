package main

import "fmt"

// Point 構造体。整数型の A と文字列型の B を持つ。
type Point struct {
	A int
	B string
}

// String() メソッドで Point 型のカスタム文字列表現を返す。
// fmt.Stringer インターフェースを実装している。
func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B) // Sprintf はフォーマットされた文字列を生成するために使われ、出力しない
}

func main() {
	// Point 構造体のインスタンスを生成し、文字列表現を出力。
	p := &Point{100, "ABC"}
	fmt.Println(p) // "<<100, ABC>>" と出力される
}
