package main

import (
	"fmt"
	"net/http"
)

// 簡単なHTTPサーバーの実装

// クライアントからのリクエストを処理
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストメソッドを確認
	fmt.Println("Method:", r.Method)
	// クライアントにメッセージを送信
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// /helloにリクエストが来たらhelloHandlerを呼び出す
	http.HandleFunc("/hello", helloHandler)

	// サーバーを起動して、ポート8080でリクエストを待機
	fmt.Println("Starting server at port 8080...")
	// HTTPサーバーを指定したポートで起動し、リクエストを受け付ける関数
	if err := http.ListenAndServe(":8080", nil); err != nil {
		// サーバー起動中にエラーがあった場合に表示
		panic(err)
	}
}

// http.ResponseWriter は、サーバーからクライアントにデータを送り返すための役割
// http.Request は、クライアントからサーバーに送られてくるリクエスト情報をまとめて持っているもの

// サーバー起動後：
// ブラウザやPostmanなどで、http://localhost:8080/hello にアクセスすると、「Hello, World!」というメッセージが表示される

/*
	　デフォルトのハンドラ:
	・第一引数にはポート番号を :8080 のように指定
	・第二引数はハンドラを指定するためのものですが、nil を渡すとデフォルトのハンドラが使用される
	・今回は、http.HandleFunc で設定したルーティングが使われる
	・/unknown のように設定されていないパスにアクセスすると、デフォルトハンドラは404 Not Foundのレスポンスを返す
*/
