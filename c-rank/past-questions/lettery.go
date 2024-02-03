// 宝くじ (paizaランク C 相当)

// 下記の問題をプログラミングしてみよう！
// 今年もパイザ宝くじの季節がやってきました。パイザ宝くじ券には 100000 以上 199999 以下の 6 桁の
// 番号がついています。毎年1つ当選番号 (100000 以上 199999 以下)が発表され、当たりクジの番号が以下のように決まります。

// 1等：当選番号と一致する番号
// 前後賞：当選番号の ±1 の番号（当選番号が 100000 または 199999 の場合，前後賞は一つしかありません）
// 2等：当選番号と下 4 桁が一致する番号（1等に該当する番号を除く）
// 3等：当選番号と下 3 桁が一致する番号（1等および2等に該当する番号を除く）

// たとえば、当選番号が 142358 の場合、当たりの番号は以下のようになります。

// 1等：142358
// 前後賞：142357 と 142359
// 2等：102358, 112358, 122358, …, 192358 （全 9 個）
// 3等：100358, 101358, 102358, …, 199358 （全 90 個）

// あなたが購入した n 枚の宝くじ券の各番号が入力されるので、それぞれの番号について、何等に当選したかを出力するプログラムを書いて下さい。

// 入力される値
// 入力は以下のフォーマットで与えられます。

// b
// n
// a_1
// a_2
// :
// a_n

// ここで、b は当選番号、n は購入した宝くじの数、a_1,…,a_n は購入した宝くじ券の番号をそれぞれ表します。


// 入力値最終行の末尾に改行が１つ入ります。
// 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
// 期待する出力
// 期待する出力は n 行から成ります。各 i 行目 (1 ≦ i ≦ n) には、a_i が何等に当たったかに応じて、以下の文字列を出力して下さい。

// 1等の場合: first
// 前後賞の場合: adjacent
// 2等の場合: second
// 3等の場合: third
// それ以外（外れ）の場合: blank

// 最後は改行し、余計な文字、空行を含んではいけません。

// 条件
// すべてのテストケースにおいて、入力される値は全て整数であり、以下の条件をみたします。

// 100,000 ≦ b ≦ 199,999
// 1 ≦ n ≦ 100
// 100,000 ≦ a_i ≦ 199,999 (1 ≦ i ≦ n)

// 入力例1
// 142358 当選番号
// 3 くじの枚数
// 195283 買った宝くじの番号
// 167358 買った宝くじの番号
// 142359 買った宝くじの番号

// 出力例1
// blank
// third
// adjacent

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main {

// 	var N int
 
// 	fmt.Scan(&N)
// 	// 標準入力された値を整数値に変える関数

// 	scanner := bufio.NewScanner(os.Stdin)

// 	var boughtLotterynumbers [] int

// 	succeceNumber := scanner.Text() 
// 	fmt.Println(succeceNumber)
// 	// 当選番号を取得

// 	boughtLottelys := scanner.Text() 
// 	fmt.Println(boughtLottelys)
// 	// 宝くじの枚数を取得

// 	// 枚数分forで展開して配列boughtLotterynumbersに値を入れる
// 	for i := 0; i < boughtLottelys && scanner.Scan(); boughtLottelys++ {
// 		fmt.Println(scanner.Text())
// 	}

// 	//if文あたりで当選したか否かの判定を行うas

// }
// ! 解けなかった問題のワイの回答⇧

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var winningNumber int
	fmt.Scan(&winningNumber)

	var n int
	fmt.Scan(&n)


	tickets := make([]int, n)


	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n && scanner.Scan(); i++ {
		ticket, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		tickets[i] = ticket
	}


	for _, ticket := range tickets {
		result := checkResult(winningNumber, ticket)
		fmt.Println(result)
	}
}

// 当選結果を確認する関数
func checkResult(winningNumber, ticket int) string {
	if winningNumber == ticket {
		return "first"
	} else if winningNumber-1 == ticket || winningNumber+1 == ticket {
		return "adjacent"
	} else if winningNumber%10000 == ticket%10000 {
		return "second"
	} else if winningNumber%1000 == ticket%1000 {
		return "third"
	}
	return "blank"
}

// 気になるメソッド


