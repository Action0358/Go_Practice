package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// map の定義（キーは文字列型、値は Person 構造体）
	people := make(map[string]Person)

	// map にデータを追加
	// people["alice"] に Name: "Alice", Age: 30 の Person 構造体を格納
	people["alice"] = Person{Name: "Alice", Age: 30}
	// people["bob"] に Name: "Bob", Age: 25 の Person 構造体を格納
	people["bob"] = Person{Name: "Bob", Age: 25}

	// key には "alice" や "bob" などの文字列キーが入る
	// person には対応する Person 構造体が格納される
	for key, person := range people {
		// key と Person の Name と Age を表示
		fmt.Printf("Key: %s, Name: %s, Age: %d\n", key, person.Name, person.Age)
	}
}
