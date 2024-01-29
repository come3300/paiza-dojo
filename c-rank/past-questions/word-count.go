//** 単語のカウント (paizaランク C 相当)

// 下記の問題をプログラミングしてみよう！
// スペースで区切られた英単語列が与えられます。
// 英単語列に含まれる英単語の出現回数を出現した順番に出力してください。

// 入力される値
// 半角スペースで区切られた長さNの文字列


// 入力値最終行の末尾に改行が１つ入ります。
// 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
// 期待する出力
// 単語、半角スペース、出現回数の順で１行に１単語で出現したすべての単語を、列に出現する順に出力してください。

// 条件
// 全てのテストケースにおいて以下の条件を満たします。

// 1 ≦ N ≦ 1,000

// 入力例1
// red green blue blue green blue

// 出力例1
// red 1
// green 2
// blue 3

// 入力例2
// Apple Apricot Orange Cherry Apple Orange Cherry Orange

// 出力例2
// Apple 2
// Apricot 1
// Orange 3
// Cherry 2
package main

import (
        "fmt"
        "strings"
)

func main() {
	// 処理対象（英単語が半角スペース区切りで並んだ文字列）
	str := "tokyo kyoto fukuoka tokyo fukuoka sapporo tokyo"

	// 半角スペースを区切り文字として、strから文字列のスライスを作成
	strSlice := strings.Split(str, " ")
	fmt.Println("strSlice:", strSlice)
	// strSlice: [tokyo kyoto fukuoka tokyo fukuoka sapporo tokyo]
	

	// 文字列の出現数をカウントするマップを作成
	counts := make(map[string]int)
    fmt.Println("counts:", counts)
	// counts: map[]


	// 文字列のスライスから要素の文字列を順に取り出し、カウンタをインクリメント
	for _, s := range strSlice {
		counts[s]++
	}
	fmt.Println("counts after loop:", counts)
	// counts after loop: map[fukuoka:2 kyoto:1 sapporo:1 tokyo:3]
	

	// 結果の出力
	for s, n := range counts {
		fmt.Printf("%s: %d\n", s, n)
	}
	
}



