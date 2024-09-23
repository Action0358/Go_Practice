package main

import (
	"fmt"
	"math"
)

func main() {
	// 絶対値の取得
	a := -9.5
	absValue := math.Abs(a)
	fmt.Printf("Absolute value of %.2f is %.2f\n", a, absValue)

	// 平方根の計算
	number := 16.0
	sqrtValue := math.Sqrt(number)
	fmt.Printf("Square root of %.2f is %.2f\n", number, sqrtValue)

	// 最大値と最小値の取得
	x := 7.5
	y := 12.3
	maxValue := math.Max(x, y)
	minValue := math.Min(x, y)
	fmt.Printf("Max of %.2f and %.2f is %.2f\n", x, y, maxValue)
	fmt.Printf("Min of %.2f and %.2f is %.2f\n", x, y, minValue)

	// 累乗の計算
	base := 2.0
	exponent := 3.0
	powerValue := math.Pow(base, exponent)
	fmt.Printf("%.2f to the power of %.2f is %.2f\n", base, exponent, powerValue)

	// 角度のラジアン変換 (角度→ラジアン)
	degrees := 90.0
	radians := degrees * (math.Pi / 180)
	fmt.Printf("%.2f degrees is %.2f radians\n", degrees, radians)

	// 正弦の計算 (ラジアンで指定)
	sineValue := math.Sin(radians)
	fmt.Printf("Sine of %.2f radians is %.2f\n", radians, sineValue)
}

// 数学的な関数や定数を提供する標準ライブラリ

// %f：浮動小数点数（float32 または float64）を表示する指定子
// .2：小数点以下の桁数を指定する部分であり、小数点以下2桁を指す
