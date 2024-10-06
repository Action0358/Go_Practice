package main

import (
	"crypto/aes"    // AES (Advanced Encryption Standard) を使った暗号化・復号化を行うためのパッケージ
	"crypto/cipher" // 暗号ブロックを使って、より高レベルな暗号化モード（例: GCM, CBC, CFB など）を実現するためのパッケージ
	"crypto/rand"   // 暗号に安全なランダムなデータを生成するためのパッケージ
	"encoding/hex"  // データを16進数（hex）表現にエンコード、あるいはデコードするためのパッケージ
	"fmt"           // 標準出力へのフォーマット出力を行うための基本的なパッケージ
	"io"            // データの読み書きを行うためのインターフェースを提供するパッケージ
)

// GoでのAES暗号化

func main() {
	// 暗号化するメッセージ（このメッセージが金庫に入る）
	plaintext := []byte("This is a secret message")

	// 32バイトの鍵（鍵のサイズはAESの種類によって16, 24, 32バイトが必要）
	key := []byte("thisis32bitlongpassphraseimusing")

	// AESブロックを作成（金庫の準備）
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// GCMモードを使う（特別な鍵のかけ方）
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	// ノンス（乱数ロックの番号）を作成
	nonce := make([]byte, aesGCM.NonceSize())                  // ノンス（Nonce）**という一時的な値（乱数のようなもの）を生成
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil { // ノンスをランダムな値で埋める
		panic(err)
	}

	// メッセージを暗号化（データを金庫に入れる）
	ciphertext := aesGCM.Seal(nil, nonce, plaintext, nil) // ノンスをランダムな値で埋める

	// 暗号化されたデータを16進数で表示（中身を見やすくする）
	fmt.Printf("Encrypted (hex): %s\n", hex.EncodeToString(ciphertext))

	// 復号化（金庫からメッセージを取り出す）
	decryptedText, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	// 復号化したメッセージを表示
	fmt.Printf("Decrypted: %s\n", decryptedText)
}

/*
	•	AESは「金庫そのもの」
	•	cipherは「金庫の鍵のかけ方」
	•	randは「金庫のランダムなロック番号」
	•	hexは「金庫の中身を見やすくする変換ツール」
	•	fmtは「金庫に何が入っているかを報告するメッセージ」
	•	ioは「金庫にデータを入れるための道具」
*/
