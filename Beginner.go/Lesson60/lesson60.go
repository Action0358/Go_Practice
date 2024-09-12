package main

import "fmt"

// Person 構造体
type Person struct {
	Name string
	Age  int
}

// NewPerson コンストラクタ関数
// これは新しい Person のインスタンスを作成して返す関数
func NewPerson(name string, age int) Person {
	return Person{ // Person{...} という記述は、Person 構造体の新しいインスタンスを作成
		Name: name, // Name フィールドに、引数 name の値を代入
		Age:  age,  // Age フィールドに、引数 age の値を代入
	}
}

func main() {
	// NewPerson コンストラクタ関数を使用して Person インスタンスを作成
	p1 := NewPerson("Alice", 25)
	p2 := NewPerson("Bob", 30)

	// Person のフィールドにアクセスして表示
	fmt.Printf("Name: %s, Age: %d\n", p1.Name, p1.Age)
	fmt.Printf("Name: %s, Age: %d\n", p2.Name, p2.Age)
}
