//** 検索履歴 (paizaランク C 相当)

// 下記の問題をプログラミングしてみよう！
// あなたが利用しているブラウザでは検索ワードの履歴を見ることができません。あなたは検索ワードの履歴を見られないのは不便だと思ったので、検索ワードの履歴を見る機能を自分でつくることにしました。

// 検索ワードの履歴とは次のようにつくられます。

// 検索ワード W が以前に入力されたことがある場合：
// 履歴中の W を削除する。
// 履歴の先頭に W を追加する。
// 検索ワード W が以前に入力されたことがない場合：
// 履歴の先頭に W を追加する。

// 検索ワード W が N 個与えられるので、N 個の検索ワードが与えられた後の履歴を表示するプログラムを書いてください。

// 入力される値
// 入力は以下のフォーマットで与えられます。

// N
// W_1
// W_2
// ...
// W_N

// 1 行目には検索ワードの数を表す整数 N が与えられます。
// 続く N 行では検索ワード W_i が与えられます。
// 続く N 行のうちの i 行目 (1 ≦ i ≦ N) には、検索ワード W_i が与えられます。検索ワード W_i は小文字のアルファベット a ~ z のみからなる文字列です。
// 入力は合計 N + 1 行であり、 最終行の末尾に改行が 1 つ入ります。

// 入力値最終行の末尾に改行が１つ入ります。
// 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
// 期待する出力
// 検索ワードを N 個入力した後の検索履歴を出力してください。
// 出力の最後に改行を入れ、余計な文字、空行を含んではいけません。

// 条件
// すべてのテストケースにおいて、以下の条件をみたします。

// 1 ≦ N ≦ 100
// 各 W_i (1 ≦ i ≦ N) に対して、W_iの文字数が20を超えない。

// 入力例1
// 5
// book
// candy
// apple
// book
// candy

// 出力例1
// candy
// book
// apple

// 入力例2
// 6
// apple
// book
// information
// note
// pen
// pineapple

// 出力例2
// pineapple
// pen
// note
// information
// book
// apple

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 標準入力からNの値を読み取る
	var N int

	// 標準入力から読み込んだデータを
	// 整数値に変える関数
	fmt.Scan(&N)

	var inputStrings []string

	// 標準入力から各行の文字列を読み取り、出力する
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < N && scanner.Scan(); i++ {

		inputString := scanner.Text()
		fmt.Println("読みよった文字列", inputString)

		// 配列に追加
		inputStrings = append(inputStrings, inputString)
	}

	// エラーがあれば出力
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "標準入力の読み込みエラー:", err)
		return
	}

	fmt.Println("代入された文字列:", inputStrings)
	// ここまでで、標準入力からの文字列を配列に代入できた。

	// 重複を削除した配列を作成

	m := make(map[string]bool)
	uniq := [] string{}

	for _, ele := range inputStrings {
		if !m[ele] {
			m[ele] = true
			uniq = append(uniq, ele)
		}
	}

	// 重複削除した配列を出力

	//! fmt.Println(uniq)   これだと配列ごと出力されちゃう
	
    for _, word := range uniq {
    fmt.Println(word)
}




}

// 処理対象（英単語が半角スペース区切りで並んだ文字列）
// str := "tokyo kyoto fukuoka tokyo fukuoka sapporo tokyo"

// 半角スペースを区切り文字として、strから文字列のスライスを作成
// strSlice := strings.Split(str, " ")
// fmt.Println("strSlice:", strSlice)
// strSlice: [tokyo kyoto fukuoka tokyo fukuoka sapporo tokyo]

// 文字列の出現数をカウントするマップを作成
// counts := make(map[string]int)
// fmt.Println("counts:", counts)
// counts: map[]

// 文字列のスライスから要素の文字列を順に取り出し、カウンタをインクリメント
// for _, s := range strSlice {
// 	counts[s]++
// }
// fmt.Println("counts after loop:", counts)
// counts after loop: map[fukuoka:2 kyoto:1 sapporo:1 tokyo:3]

// 結果の出力
// for s, n := range counts {
// 	fmt.Printf("%s: %d\n", s, n)
// }
