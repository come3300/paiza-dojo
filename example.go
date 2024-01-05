
//* PaizaのPaizaの値取得・出力方法 解説

// Paizaでの入力される値を
// 変数に代入して処理をする方法についての解説である。

//** 入力される値
// 2
// 2 5
// 3 4

// このサンプルの値をhello=,world= に代入して
// そこから問題の処理を行なっていくための処理

// ChatGPTに聞くと.txtなどのファイルに値を入れて
// その値を取り込んで処理をするような記述になるがそうではなく
// あくまで入力された値（入力される値のサンプルのような値）を
// 取り込んで実行しなければならない

// paiza.ioの入力欄に入力して正しく出力結果が得られるようにしなければならない
// 下記リンクのpaiza.ioを参照
// https://paiza.io/ja/projects/new

//** Paizaでの入力例

//** サンプルコード
//** 入力される値
// 2
// 2 5
// 3 4
// このテストケースでは、最初の値は、その後入力される行数を示す(2行の入力がある)
// 2行目以降は、helloとworldの値が[,]区切りで書かれています。

// 期待する出力
// hello = 2 , world = 5
// hello = 3 , world = 4

package main
import "fmt"
import "bufio"
import "os"
import "strconv"
import "strings"

// main 関数がプログラムのエントリーポイント
func main() {
	// 標準入力からデータを読み取る Scanner を作成
	sc := bufio.NewScanner(os.Stdin)

	// 1行目の入力を読み取り、整数に変換
	sc.Scan()
	var n, _ = strconv.Atoi(sc.Text())

	// n 回だけ以下の処理を繰り返す
	for i := 0; i < n; i++ {
		// 再び標準入力から行を読み取り
		sc.Scan()
		// 行をスペースで分割し、文字列のスライスにする
		var s = strings.Split(sc.Text(), " ")
		// 文字列の結合と出力
		fmt.Println("hello = " + s[0] + " , world = " + s[1])
	}
}
