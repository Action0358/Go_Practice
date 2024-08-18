package main

import "fmt"

type User struct {
	name string
}

func (b User) cal(weight, height float64) (result float64) {
	result = ((weight / height) / weight) * 10000
	return
}

func main() {
	user := User{"Daiyu"}
	fmt.Println(user.cal(60, 172))
}
