// 入力される値
// 3つの文字列がカンマ区切りで1行で与えれます。


// 入力値最終行の末尾に改行が１つ入ります。
// 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
// 期待する出力
// 3行での出力

// 条件
// すべてのテストケースにおいて、以下の条件をみたします。

// ・入力される各文字列は1文字以上100文字以下の文字列
// ・入力される各文字列の各文字は英小文字または大文字または数字

// 入力例1
// aaaaa,bbbbbb,cccc

// 出力例1
// aaaaa
// bbbbbb
// cccc

// 入力例2
// a,b,c

// 出力例2
// a
// b
// c

// 入力例3
// a,a,a

// 出力例3
// a
// a
// a

// 入力例4
// 10,31415926,a

// 出力例4
// 10
// 31415926
// a

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 標準入力から1行の文字列を読み取る
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// カンマで区切られた文字列を配列に分割
	strArray := strings.Split(input, ",")

	// 配列の要素を新しい行に出力
	for _, str := range strArray {
		fmt.Println(str)
	}
}
