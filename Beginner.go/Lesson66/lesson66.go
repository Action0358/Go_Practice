package main

import "fmt"

// Person 構造体
type Person struct {
	Name string
	Age  int
}

// ToString メソッドを Person 構造体に実装
func (p *Person) ToString() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age) // %s は文字列に、%d は整数を表すフォーマット指定子
}

// Car 構造体
type Car struct {
	Number string
	Model  string
}

// ToString メソッドを Car 構造体に実装
func (c *Car) ToString() string {
	return fmt.Sprintf("Number: %s, Model: %s", c.Number, c.Model)
}

// Stringify インターフェース
type Stringify interface {
	ToString() string
}

func main() {
	// Stringify インターフェースを実装した構造体のスライス
	entities := []Stringify{ // スライスの宣言
		// & はアドレス演算子で、変数のメモリアドレスを取得し、そのポインタを作成
		&Person{Name: "Taro", Age: 21},
		&Car{Number: "123-456", Model: "AB-1234"},
	}

	// 各エンティティの ToString メソッドを呼び出して表示
	for _, entity := range entities { // _ = インデックスなし, entity = 値あり
		fmt.Println(entity.ToString())
	}
}
