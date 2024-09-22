package main

import "fmt"

// Point 構造体。整数型の A と文字列型の B を持つ。
type Point struct {
	A int
	B string
}

// String() メソッドで Point 型のカスタム文字列表現を返す。
func (p *Point) String() string { // レシーバによって main 関数の p がメソッドを呼び出せるようになっている
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B) // Sprintf はフォーマットされた文字列を生成するために使われ、出力しない
}

func main() {
	// Point 構造体のインスタンスを生成し、文字列表現を出力。
	p := &Point{100, "ABC"}
	fmt.Println(p) // "<<100, ABC>>" と出力される
}

// p は Point 型のポインタを指す変数（12, 18行目）
// *Point はそのポインタが指している元の Point 構造体のデータを操作するための型指定（19行目のデータ）
// p := &Point{100, "ABC"} は Point 構造体を生成し、そのアドレス（ポインタ）の取得を行う
// fmt.Println 関数が p に対して String() メソッドを呼び出して、main関数内でフォーマットされた文字列が出力される
