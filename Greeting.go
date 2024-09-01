package main // mainパッゲージを記述

import "fmt" //fmtパッケージの中にある関数などが使えるようになる

func main() { //main関数の定義、{}内にプログラムの処理を記述
	fmt.Println("Good morning") // Println関数を使ってGood morningを表示
	fmt.Println("Good afternoon")

	fmt.Println("Good\nevening") // \nは改行を意味する
}
