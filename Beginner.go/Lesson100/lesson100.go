package main

import (
	"fmt"
	"net/http"
	"os"
)

// HTMLファイルをサーバーから返す

func serveHTML(w http.ResponseWriter, r *http.Request) {
	// HTMLファイルのパス
	htmlFile := "index.html"

	// ファイルを開く
	file, err := os.Open(htmlFile)
	if err != nil {
		// エラーが発生した場合、500エラーを返す
		http.Error(w, "File not found.", 500)
		return
	}
	defer file.Close()

	// HTMLファイルをクライアントに送信
	// ファイルを一行ずつ読み込んで送信する
	fmt.Fprintln(w, "Content-Type: text/html")
	http.ServeFile(w, r, htmlFile)
}

func main() {
	// "/index"パスにアクセスされたらserveHTML関数が実行されるように設定
	http.HandleFunc("/index", serveHTML)

	// サーバーを8080ポートで開始
	fmt.Println("サーバーを開始します。http://localhost:8080/index") // index.html の内容が表示される
	http.ListenAndServe(":8080", nil)
}

/*
	処理の流れ
	1. ブラウザから /index にリクエストが送られる
	2. サーバーは serveHTML 関数を呼び出す
	3. サーバーは指定された index.html ファイルを開き、内容をクライアント（ブラウザ）に送信する
	4. ブラウザは受け取ったHTMLを表示する
*/
