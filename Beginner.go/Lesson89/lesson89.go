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
	// 暗号化したいデータ
	plaintext := []byte("This is a secret message")

	// 暗号化・復号化に使う32バイトのキー（AES-256）
	key := []byte("thisis32bitlongpassphraseimusing")

	// AES-GCMモードでの暗号化を行う
	encrypted, err := encryptAESGCM(plaintext, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encrypted (hex): %s\n", hex.EncodeToString(encrypted))

	// 暗号化されたデータを復号化する
	decrypted, err := decryptAESGCM(encrypted, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypted: %s\n", decrypted)
}

// AES-GCMモードで暗号化
func encryptAESGCM(plaintext, key []byte) ([]byte, error) {
	// AESブロックを作成
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// GCMモードの作成
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// ノンス（nonce）の生成
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// 暗号化
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

// AES-GCMモードで復号化
func decryptAESGCM(ciphertext, key []byte) ([]byte, error) {
	// AESブロックを作成
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// GCMモードの作成
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// ノンスサイズを取得
	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// 復号化
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
