package main

import (
	"bufio"
	"fmt"
	"os"
)

// 入力される値
// 文字列A, B, Cが各行で入力されます。


// 入力値最終行の末尾に改行が１つ入ります。
// 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
// 期待する出力
// 3行での出力

// 条件
// すべてのテストケースにおいて、以下の条件をみたします。

// ・A, B, Cは1文字以上100文字以下の文字列
// ・A, B, Cの各文字は英小文字または大文字または数字または半角スペース

// 入力例1
// aaaaa
// bbbbbb
// cccc

// 出力例1
// aaaaa
// bbbbbb
// cccc

// 入力例2
// a b c
// d e f
// g h i

// 出力例2
// a b c
// d e f
// g h i

// 入力例3
// 2000 500
// aba 200 3a 3
// ppppppppp

// 出力例3
// 2000 500
// aba 200 3a 3
// ppppppppp

// 入力例4
// 0x04a12bE avacc
// fffew2feafewafez
// asas111 abbaewafew

// 出力例4
// 0x04a12bE avacc
// fffew2feafewafez
// asas111 abbaewafew

func main() {
	// 標準入力から読み込むためのScannerを作成
	scanner := bufio.NewScanner(os.Stdin)

	// 3行の文字列を格納するための変数を用意 こんな風に3つ一気に定義できるのか　
	var A, B, C string

	// 1行ずつ入力を受け取る
	// ScannerのScanメソッドは入力がある限りtrueを返し、エラーまたはEOF（入力の終了）の場合はfalseを返す
	for i := 0; i < 3 && scanner.Scan(); i++ {
		// 入力を変数に格納
		switch i {
		case 0:
			A = scanner.Text()
		case 1:
			B = scanner.Text()
		case 2:
			C = scanner.Text()
		}
	}

	// エラーがあれば出力
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "標準入力の読み込みエラー:", err)
		return
	}

	// 結果を標準出力に出力
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
