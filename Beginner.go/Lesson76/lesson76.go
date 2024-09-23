// rand パッケージは、Goで疑似乱数を生成するためのパッケージ

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// シード値を現在の時間から設定 (これで実行するたびに異なる乱数が生成される)
	rand.Seed(time.Now().UnixNano())

	// ランダムな整数 (0〜99)
	randomInt := rand.Intn(100)
	fmt.Printf("ランダムな整数: %d\n", randomInt)

	// ランダムな浮動小数点数 (0.0〜1.0)
	randomFloat := rand.Float64()
	fmt.Printf("ランダムな浮動小数点数: %.4f\n", randomFloat)

	// 特定範囲のランダムな整数 (50〜100)
	randomRangeInt := rand.Intn(51) + 50
	fmt.Printf("50から100の間のランダムな整数: %d\n", randomRangeInt)
}
