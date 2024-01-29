// 【次元配列の入出力】i番目の出力 Boss (paizaランク D 相当)

// 入力される値
// 1 行目に整数 N と整数 K が与えられます。
// 2 行目に N 個の整数 a_i が半角スペース区切りで与えられます。
// 以下の形式で標準入力によって与えられます。


// N K
// a_1 ... a_N

// 入力値最終行の末尾に改行が１つ入ります。
// 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
// 期待する出力
// K 番目の整数 a_K を出力してください。
// また、末尾に改行を入れ、余計な文字、空行を含んではいけません。


// a_K
// 条件
// すべてのテストケースにおいて、以下の条件をみたします。

// * N は 1 以上 100 以下の整数
// * K は 1 以上 N 以下の整数
// * a_i は 1 以上 100 以下の整数

// 入力例1
// 8 3
// 3 1 8 1 3 8 8 1

// 出力例1
// 8

//! 解説

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 標準入力から1行取得
	scanner := bufio.NewScanner(os.Stdin)
//! bufio.NewScanner(os.Stdin):
//! bufio.NewScanner メソッドは os.Stdin からの標準入力を読み込む Scanner を生成します。
//! Scanner はテキストデータを逐次的に読み込むための機能を提供します。

	scanner.Scan()
//! scanner.Scan():
//! scanner.Scan メソッドは次のトークン（行）を読み込む操作を実行します。読み込みが成功すると、
//! scanner.Text() メソッドで読み込んだテキストを取得できます。

	line1 := scanner.Text() // scanner.Scan()の文字列を取得   ここまでで1行目をline1に格納

	// 1行目を空白で分割してNとKに分ける
	values := strings.Fields(line1)
//! strings.Fields(line1):
//! strings.Fields メソッドは与えられた文字列を空白文字で分割し、それぞれのフィールド（単語）を含む文字列のスライスを返します。
//!例えば、もし line1 が "5 3" だった場合、strings.Fields(line1) は ["5", "3"] のような文字列のスライスを生成し、これが values に代入されます。

	// Nの値を取得したが、使用しないので _ で無視
	_, _ = strconv.Atoi(values[0]) // 0番目の配列の値を文字列から整数に変換
	// _（アンダースコア）:
//! アンダースコア _ は「空白識別子」または「無視識別子」と呼ばれ、
//! 値を捨てるために使用されます。この場合、Atoi メソッドから返された値を無視します。
//! strconv.Atoi(values[0]):
//! strconv.Atoi メソッドは文字列を整数に変換します。この場合、文字列 values[0] を整数に変換しています。

	K, _ := strconv.Atoi(values[1]) //! 同様に、文字列 values[1] を整数に変換し、K にその値を代入しています。
//! K, _ := strconv.Atoi(values[1]):


	// 2行目を整数のスライスに変換
	scanner.Scan()
	line2 := scanner.Text()
	numbers := parseIntArray(line2)
//!parseIntArray(line2):
//! parseIntArray 関数は、与えられた文字列を整数のスライスに変換するための自作のヘルパー関数です。

	// K番目の整数を出力
	result := numbers[K-1]
	fmt.Println(result)
//! fmt.Println(result):
//! fmt.Println メソッドは標準出力に引数で渡された値を表示します。この場合、result の値が表示されます。

}

// スペース区切りの文字列を整数のスライスに変換する関数
func parseIntArray(input string) []int {
	values := strings.Fields(input)
	var numbers []int
	for _, value := range values {
		num, _ := strconv.Atoi(value)
		numbers = append(numbers, num)
	}
	return numbers
}









