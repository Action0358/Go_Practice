package main

import "fmt"

// Stringer インターフェースは標準ライブラリ
/*
	type Stringer interface {
		String() string
	}
*/

// Season 型は int を基にしたカスタム型
type Season int

// const は定数の定義であり、プログラムが実行される間に変更されることがない
// iota を使用して、Spring から Winter に連続した整数 (0, 1, 2, 3) を割り当てる
const (
	Spring Season = iota // Spring = 0
	Summer               // Summer = 1
	Autumn               // Autumn = 2
	Winter               // Winter = 3
)

// Stringer インターフェースを実装する
// fmt.Stringer インターフェースは String() メソッドを持ち、このメソッドがオブジェクトの文字列表現を返すように定義する
// ここでは Season 型に対して文字列を対応させている
func (s Season) String() string {
	// Season 型の値に応じた文字列を返す。
	switch s {
	case Spring:
		return "Spring" // Season = 0 の場合、"Spring" を返す
	case Summer:
		return "Summer" // Season = 1 の場合、"Summer" を返す
	case Autumn:
		return "Autumn" // Season = 2 の場合、"Autumn" を返す
	case Winter:
		return "Winter" // Season = 3 の場合、"Winter" を返す
	default:
		return "Unknown" // それ以外の値が来た場合は "Unknown" を返す
	}
}

func main() {
	// 定義したSeasonを表示
	fmt.Println(Spring) // "Spring" と表示される
	fmt.Println(Summer) // "Summer" と表示される
	fmt.Println(Autumn) // "Autumn" と表示される
	fmt.Println(Winter) // "Winter" と表示される
}
