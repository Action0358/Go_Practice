package main

import "fmt"

// Person 構造体の定義
type Person struct {
	Name string
	Age  int
}

func main() {
	// Person のスライスを作成
	// スライスは可変長の配列のようなデータ構造
	var people []Person

	// スライスに Person を追加
	people = append(people, Person{Name: "Alice", Age: 30})
	people = append(people, Person{Name: "Bob", Age: 25})
	people = append(people, Person{Name: "Charlie", Age: 35})

	// スライス内の全ての Person をループして表示
	for _, person := range people {
		fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
	}
}
