package main

import "fmt"

type User struct {
	gender string
	age    int
}

func main() {
	u := User{"male", 20}

	fmt.Println(u)
}
