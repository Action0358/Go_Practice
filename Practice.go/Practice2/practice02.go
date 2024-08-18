package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// 別の記述1
// func main() {
// 	i := 1
// 	for i <= 10 {
// 		fmt.Println(i)
// 		i++
// 	}
// }

// 別の記述2
// func main() {
//		i := 1
//		for {
//			fmt.Println(i)
// 			if i == 10 {
// 				break
// 			}
//			i++
// 		}
// }
