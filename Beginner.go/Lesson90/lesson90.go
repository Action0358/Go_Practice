package main

import (
	"encoding/json" // JSONのエンコードとデコードを行うためのパッケージ
	"fmt"           // 標準出力にメッセージを表示するためのパッケージ
)

// Userという構造体を定義（JSONデータの構造に対応）
type User struct {
	Name  string `json:"name"`  // JSONの"名前"フィールドにマッピング
	Email string `json:"email"` // JSONの"メール"フィールドにマッピング
	Age   int    `json:"age"`   // JSONの"年齢"フィールドにマッピング
}

func main() {
	// User構造体のインスタンスを作成
	user := User{
		Name:  "Alice",
		Email: "alice@example.com",
		Age:   30,
	}

	// UserインスタンスをJSONにエンコード（変換）
	jsonData, err := json.Marshal(user)
	if err != nil {
		// エンコード中にエラーが発生した場合はメッセージを表示
		panic(err)
	}

	// エンコードされたJSONデータを表示
	fmt.Printf("JSON Data: %s\n", jsonData)

	// 新たに用意したJSONデータをデコードして、Goの構造体に復元する
	jsonString := `{"name":"Bob","email":"bob@example.com","age":25}`

	// JSON文字列をUser構造体にデコード（復元）
	var newUser User
	if err := json.Unmarshal([]byte(jsonString), &newUser); err != nil { // &newUserを渡していることで、newUser（フィールド）にデータが書き込まれる
		// デコード中にエラーが発生した場合はメッセージを表示
		panic(err)
	}

	// デコードされたUser構造体の情報を表示
	fmt.Printf("Decoded User: %+v\n", newUser)
}
