// context パッケージは、Goでタイムアウトやキャンセル機能を持つ操作

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// タイムアウトを5秒に設定したContextを作成
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // 関数が終了する際にキャンセルされる

	// タスクを実行
	done := make(chan string)
	go func() {
		// 3秒後に完了する処理をシミュレート
		time.Sleep(3 * time.Second)
		done <- "タスクが完了しました！"
	}()

	// タスクが完了するか、タイムアウトするかを待機
	select {
	case result := <-done:
		fmt.Println(result)
	case <-ctx.Done():
		fmt.Println("タイムアウトしました！")
	}
}

/*
	1.このコードは最初に5秒のタイムアウトを設定
	2.ゴルーチンで3秒間かかる処理をシミュレートし、完了後に done チャネルにメッセージを送る
	3.メインの select 文は、タスクが完了するかタイムアウトするかを待つ
	4.この場合、タスクは3秒で完了するため、"タスクが完了しました！" が表示される
	5.しタスクが6秒以上かかっていた場合は、タイムアウトして "タイムアウトしました！" が表示される
*/
