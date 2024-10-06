package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// JSONデータをPOSTリクエストで送信

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// 送信するデータを作成（User構造体）
	user := User{
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}

	// データをJSONにエンコード
	jsonData, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	// POSTリクエストを作成
	// 最初の引数は、リクエストを送信する URL
	// 2番目の引数は、コンテンツタイプ
	// 3番目の引数は、エンコードされたJSONデータをバッファに変換して送信
	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		// エラーがあれば表示して終了
		panic(err)
	}
	defer resp.Body.Close()

	// サーバーのレスポンスを表示
	fmt.Println("ステータスコード:", resp.Status)
}
