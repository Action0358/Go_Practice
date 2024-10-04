package main

import (
	"fmt"
	"sync"
)

// sync.Mutex の役割

var balance int = 10000 // 口座の残高
var mu sync.Mutex       // ロック用のMutex

func withdraw(amount int, wg *sync.WaitGroup) {
	defer wg.Done() // この関数が終わったら、WaitGroupに「作業が終わった」と伝える

	mu.Lock() // 口座への操作をロック
	if balance >= amount {
		fmt.Printf("Withdrawing %d\n", amount)
		balance -= amount
	} else {
		fmt.Println("Insufficient funds") // 資金不足
	}
	mu.Unlock() // 操作が終わったらロックを解除
}

func main() {
	var wg sync.WaitGroup //  WaitGroupを作成（ゴルーチンを追跡する）

	wg.Add(2)
	go withdraw(5000, &wg) // 1人目の引き出し
	go withdraw(5000, &wg) // 2人目の引き出し

	wg.Wait()
	fmt.Printf("Final balance: %d\n", balance)
}

/*
    1.	初期残高は10,000円
	2.	2つのゴルーチンが同時に5000円の引き出しを試みる
	3.	mu.Lock() で「ロック」をかけるので、一度に1つのゴルーチンしか残高にアクセスできない
	4.	最初にアクセスできたゴルーチンが5000円を引き出し、その後ロックを解除する
	5.	次のゴルーチンがアクセスした時、残高が5000円に減っているので、もう5000円引き出すことができる
	6.	最終的に、残高は0円になる
*/
