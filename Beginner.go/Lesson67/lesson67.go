package main

import "fmt"

//　error インターフェースは標準ライブラリ
// インターフェースを実装するためには、Error()メソッドを持つ構造体を定義する必要がある
/*
	type error interface {
		Error() string
	}
*/

// MyError 構造体: カスタムエラーの型
type MyError struct {
	Message string // エラーメッセージ
	ErrCode int    // エラーコード
}

// Error メソッド: error インターフェースを実装するためのメソッド
// エラーメッセージを返す
func (e *MyError) Error() string { // e.Message の値が string 型なので、string の戻り値型として string を定義
	return e.Message
}

// RaisError 関数: カスタムエラーを生成して返す
// 戻り値として MyError 型のポインタを返すため、MyError が error インターフェースを実装している必要がある
func RaisError() error { // error は戻り値の型
	return &MyError{Message: "カスタムエラーが発生しました", ErrCode: 1234}
}

func main() {
	// カスタムエラーを発生させて err に代入
	err := RaisError()

	// err に対し Error() メソッドより MyError の Message フィールドを返し表示する
	fmt.Println(err.Error())

	// 型アサーションで、err が MyError 型かを確認
	if e, ok := err.(*MyError); ok {
		// MyError 型ならエラーコードを表示
		fmt.Println(e.ErrCode)
	}
}

// lesson24.go 参照
// lesson27.go 参照
// lesson57.go 参照
