// 入力される値
// 1行目でNが与えらます。
// 2行目でN個の文字列がカンマ区切りで与えれます。


// 入力値最終行の末尾に改行が１つ入ります。
// 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
// 期待する出力
// N行での出力

// 条件
// すべてのテストケースにおいて、以下の条件をみたします。

// ・1 <= N <= 10
// ・入力される各文字列は1文字以上100文字以下の文字列
// ・入力される各文字列の各文字は英小文字または大文字または数字

// 入力例1
// 3
// aaaaa,bbbbbb,cccc

// 出力例1
// aaaaa
// bbbbbb
// cccc

// 入力例2
// 1
// abc

// 出力例2
// abc

// 入力例3
// 5
// 3,1,4,1,5

// 出力例3
// 3
// 1
// 4
// 1
// 5

// 入力例4
// 10
// 31415926,qqqqqq,abab313,xyz31131,5,6,7,8,9,10

// 出力例4
// 31415926
// qqqqqq
// abab313
// xyz31131
// 5
// 6
// 7
// 8
// 9
// 10


package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 標準入力からNの値を読み取る
	var N int
	fmt.Scan(&N)

	// 標準入力からN個のカンマ区切りの文字列を読み取り、配列に格納
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputLine := scanner.Text()
	strArray := strings.Split(inputLine, ",")

	// N個の文字列を新しい行に出力
	for i := 0; i < N; i++ {
		fmt.Println(strArray[i])
	}
}
