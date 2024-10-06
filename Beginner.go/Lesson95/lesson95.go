package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// HTTPリクエストのタイムアウト

func main() {
	// 2秒間のタイムアウトを設定
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// リクエストを作成
	req, err := http.NewRequestWithContext(ctx, "GET", "https://httpbin.org/delay/3", nil) // このURLはサーバー側で3秒の遅延が発生するエンドポイント
	if err != nil {
		panic(err)
	}

	// HTTPリクエストを実行
	client := &http.Client{}
	resp, err := client.Do(req)

	// タイムアウトやエラーハンドリング
	if err != nil {
		fmt.Println("リクエスト中にエラー:", err)
		return
	}
	// リソースを解放
	defer resp.Body.Close() // メモリの無駄遣いを防ぐ

	// 成功した場合の応答を表示
	fmt.Println("リクエスト成功:", resp.Status)
}

/*
	1.	2秒間のタイムアウトを設定したコンテキスト (ctx) を作成します。
	2.	GET リクエストを作成し、2秒以内に応答がない場合にタイムアウトを発生させるようにします。
	3.	リクエストを実行します。今回は、サーバー側が3秒遅らせて応答するエンドポイントにリクエストを送信しています。
	4.	タイムアウトが発生する場合は、エラーメッセージが表示されます。<- 今回は4が該当する
	5.	リクエストが成功した場合は、レスポンスのステータス（例: 200 OK）を表示します。
*/
