package main

import "fmt"

// Person 構造体
type Person struct {
	Name string
	Age  int
}

// Employee 構造体は Person 構造体を埋め込んでいる
type Employee struct {
	Person   // Person を埋め込むことで、そのフィールドやメソッドを引き継ぐ
	Position string
}

func main() {
	// Employeeのインスタンスを作成
	emp := Employee{
		Person: Person{
			Name: "John",
			Age:  30,
		},
		Position: "Engineer",
	}

	// 埋め込まれたフィールドを直接参照可能
	fmt.Printf("Name: %s, Age: %d, Position: %s\n", emp.Name, emp.Age, emp.Position)
}
