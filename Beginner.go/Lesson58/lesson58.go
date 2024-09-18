package main

import "fmt"

// Personという構造体を定義
type Person struct {
	Name string
	Age  int
}

// Personに対して自己紹介を行うIntroduceメソッドを定義
func (p Person) Introduce() {
	// Printf はフォーマット指定子を使って、データをカスタマイズして表示
	fmt.Printf("Hi, my name is %s and I am %d years old.\n", p.Name, p.Age) // %s は文字列に、%d は整数を表すフォーマット指定子
}

// Personの年齢を1つ増やすHaveBirthdayメソッドを定義
func (p *Person) HaveBirthday() {
	p.Age++
}

func main() {
	// Person構造体のインスタンスを作成
	person := Person{Name: "John", Age: 30}

	// Introduceメソッドを呼び出し
	person.Introduce()

	// HaveBirthdayメソッドを呼び出して年齢を増やす
	person.HaveBirthday()

	// 年齢が増えたことを確認
	person.Introduce()
}
