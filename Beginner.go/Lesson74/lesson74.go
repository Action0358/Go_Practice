// time パッケージの関数やメソッド

package main

import (
	"fmt"
	"time"
)

func main() {
	// 現在の時刻を取得（UTC (協定世界時)）
	now := time.Now()

	fmt.Println("現在の時刻:", now)
	fmt.Println(now.Format("2006-01-02 15:04:05")) // 2006-01-02 15:04:05 はGoが固定したフォーマット指定子

	// 1時間後
	oneHourLater := now.Add(1 * time.Hour)
	fmt.Println("1時間後:", oneHourLater)

	// 30分前
	thirtyMinutesAgo := now.Add(-30 * time.Minute)
	fmt.Println("30分前:", thirtyMinutesAgo)

	// 時間の差分
	t1 := time.Now()
	t2 := t1.Add(2 * time.Hour) // 2時間後の時刻
	duration := t2.Sub(t1)      // 時間差を求める（t2 - t1）
	fmt.Println("差分:", duration)

	// 時間のスリープ
	fmt.Println("スリープ開始")
	time.Sleep(2 * time.Second) // 2秒スリープ
	fmt.Println("スリープ終了")
}

// time.Now()：現在の時刻を取得。
// time.Sleep(duration)：指定された時間だけ一時停止。
// time.Parse(layout, value)：文字列から時刻を解析。
// time.Format(layout)：時刻を文字列にフォーマット。
// time.Add(duration)：指定した時間を加算。
