package main

import (
	"fmt"
	"sync"
	"time"
)

// sync.WaitGroupの役割

func pack(name string, wg *sync.WaitGroup) {
	defer wg.Done() // このゴルーチンの作業が終わったら、WaitGroupに「完了」を通知する
	fmt.Printf("%s is packing...\n", name)
	time.Sleep(2 * time.Second) // 荷物を詰めるのに2秒かかる
	fmt.Printf("%s finished packing!\n", name)
}

func main() {
	var wg sync.WaitGroup //  WaitGroupを作成（ゴルーチンを追跡する）

	// 3つの準備タスクがあるので、WaitGroupに3を設定
	wg.Add(3)

	// 各家族メンバーの準備を並行処理
	go pack("Mom", &wg)
	go pack("Dad", &wg)
	go pack("Kid", &wg)

	// 全員の準備が終わるまで待つ
	wg.Wait()

	fmt.Println("Everyone is ready! Let's go!")
}

/*
	1.	wg.Add(3):
	•	3つの並行処理（家族の準備）を始めるので、WaitGroupに「3人分の作業があるよ」と伝える

	2.	go pack("Mom", &wg)、go pack("Dad", &wg)、go pack("Kid", &wg):
	•	pack 関数を並行で実行します。それぞれが荷物を詰める作業を行う
	•	この関数の中で wg.Done() が呼ばれると、WaitGroupに「この作業は終わった」と伝える

	3.	wg.Wait():
	•	全員の準備が終わるまでプログラムがここで「待つ」動作をして、すべての wg.Done() が呼ばれるまで、ここで停止する

	4.	準備完了後:
	•	全員の準備が終わると wg.Wait() が終了し、「全員準備完了！出発しよう！」というメッセージが表示される
*/
