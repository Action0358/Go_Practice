package main

import "fmt"

// 既存の int 型を基に独自の型 Age を定義
type Age int

// 既存の string 型を基に独自の型 Name を定義
type Name string

// 既存の構造体 Person を定義
type Person struct {
	Name Name // 独自型 Name を使用
	Age  Age  // 独自型 Age を使用
}

func main() {
	// 独自型を使用して変数宣言
	var myAge Age = 30
	var myName Name = "Alice"

	// Person 構造体を使用してデータを作成
	person := Person{
		Name: myName,
		Age:  myAge,
	}

	// 独自型のデータを表示
	fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
}
