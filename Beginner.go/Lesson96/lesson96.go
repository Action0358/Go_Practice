// net/url パッケージは、URL（Uniform Resource Locator）を扱うための便利な機能

package main

import (
	"fmt"
	"net/url"
)

// URLの解析とクエリパラメータの取得

func main() {
	// 解析するURL
	rawURL := "https://example.com/search?query=golang&sort=asc"

	// URLを解析する
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		panic(err) // エラーがあればプログラムを終了
	}

	// スキーム（https）を表示
	fmt.Println("Scheme:", parsedURL.Scheme)

	// ホスト名（example.com）を表示
	fmt.Println("Host:", parsedURL.Host)

	// パス（/search）を表示
	fmt.Println("Path:", parsedURL.Path)

	// クエリパラメータを表示（クエリパラメータとは、URLの「？」以降に続く部分）
	fmt.Println("Query Parameters:")

	// parsedURL.Query() を使って、URLのクエリパラメータを取得
	queryParams := parsedURL.Query()
	for key, values := range queryParams {
		// 各パラメータのキーと値を表示
		fmt.Printf("  %s: %s\n", key, values[0])
	}
}
