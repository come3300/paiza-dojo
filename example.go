
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

func example() {
// 本当はmain()だがファイル構成の都合上exampleに

  sc := bufio.NewScanner(os.Stdin)
// os.Stdinからの標準入力を読み取るScannerを作成します。

  sc.Scan()
  var n, _ = strconv.Atoi(sc.Text())
//   最初の行で、入力から整数 n を読み取ります。_は無視される値で、
//   Atoi関数は整数とエラーを返しますが、ここではエラーを無視しています。


  for i := 0; i < n; i++ {
    sc.Scan()
	// 各行を読み取り、スペースで分割します。strings.Splitは、
	// 指定された区切り文字で文字列を分割し、結果を文字列のスライスとして返します。
    var s = strings.Split(sc.Text(), " ")
    fmt.Println("hello = " + s[0] + " , world = " + s[1])
	// スプリットされた文字列を連結して "hello = {s[0]} , world = {s[1]}" という形式で出力します。
  }


}

