package main

import (
	"fmt"
	"io"
	"net/http"
)

// HTTP GET リクエストの実行

func main() {
	// http.Clientの作成
	client := &http.Client{}

	// GETリクエストを作成
	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		// エラーがあれば表示して終了
		panic(err)
	}
	defer resp.Body.Close() // リソースを解放

	// レスポンスの本文を読み取る
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// サーバーから受け取ったデータを表示（body変数を文字列に変換 ）
	fmt.Println(string(body))
}

// http.Client を使うことで、外部サーバーにHTTPリクエストを送り、データを取得することができる
// これは、APIから情報を取得するのに便利な方法
