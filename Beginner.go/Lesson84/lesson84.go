// regexp パッケージは、Goで正規表現（RegEx）を扱うための標準ライブラリ

package main

import (
	"fmt"
	"regexp"
)

// パターンマッチング

func main() {
	// パターンを定義: メールアドレス形式を検証
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// 正規表現パターンをコンパイル
	re := regexp.MustCompile(pattern)

	// 検証する文字列
	email := "example@domain.com"

	// パターンにマッチするか確認
	if re.MatchString(email) {
		fmt.Println("有効なメールアドレスです")
	} else {
		fmt.Println("無効なメールアドレスです")
	}
}

// ^[a-zA-Z0-9._%+-]+: 英数字と特定の記号（.、_、% など）が1文字以上連続
// @: 必須の「@」シンボル
// [a-zA-Z0-9.-]+: ドメイン名の部分
// \.[a-zA-Z]{2,}$: ドットの後に2文字以上のアルファベット
