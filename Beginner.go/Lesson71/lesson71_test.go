package lesson71

import "testing"

// TestAdd 関数は Add 関数の動作をテストします
// テスト関数は TestXxx という命名規則に従う必要がある
// 引数 t *testing.T は、テストの結果を記録・表示するためのオブジェクト
func TestAdd(t *testing.T) {
	// Add(2, 3) の結果を result に格納
	result := Add(2, 3)

	// 期待される結果は 5
	expected := 5

	// 結果が期待と一致しなければエラーを表示
	if result != expected {
		// t.Errorf はフォーマット付きでエラーメッセージを出力し、go test 実行時にその結果が表示される
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

// 実行コマンド go test
