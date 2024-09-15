package main

import "fmt"

// 組み込みの error インターフェース
// errorは標準ライブラリに含まれており、定義しなくてもerror型を扱うことができる
/* type error interface {
    Error() string
} */

// MyError 構造体: エラーメッセージとエラーコードを保持するカスタムエラー型
type MyError struct {
	Message string // エラーメッセージ
	ErrCode int    // エラーコード
}

// eはレシーバー変数で、Error メソッドが *MyError型のインスタンスに対して実行される
// Error メソッド: MyError 構造体が error インターフェースを実装するためのメソッド(MyError 型を error として扱える)
// string 型によって、このメソッドはエラーメッセージを返します
func (e *MyError) Error() string {
	return e.Message // エラーメッセージを返す
}

// RaisError 関数: カスタムエラーを発生させる関数
// error インターフェースを返すため、エラーとして扱える
func RaisError() error {
	// &MyError によりポインタ型の MyError を返す
	return &MyError{Message: "カスタムエラーが発生しました", ErrCode: 1234}
}

func main() {
	// エラーを発生させて err という変数に代入
	err := RaisError()

	// エラーメッセージを表示 (エラーインターフェースの Error メソッドが呼ばれる)
	fmt.Println(err.Error()) // 出力: カスタムエラーが発生しました

	// 型アサーション: err が MyError 型であるかどうかをチェック
	e, ok := err.(*MyError)
	if ok {
		// err が MyError 型ならば、エラーコードを表示する
		fmt.Println(e.ErrCode) // 出力: 1234
	}
}
