package main

// エンジニアであり、社員を管理を管理する立場にあるあなたは、効率的に社員を管理するために、
// 各社員の社員番号 number と名前 name を持ち、加えて情報を返す関数を持つような構造体、
// すなわち次のようなメンバ変数とメンバ関数を持つ社員クラス class employee を作成することにしました。

// メンバ変数

// number : 整数 社員番号
// name : 文字列 名前

// メンバ関数
// getnum(){
//     return number;
// }
// getname(){
//     return name;
// }

// 入力で make number name と入力された場合は変数に number , name を持つ社員を作成し、
// getnum n と入力された場合は n 番目に作成された社員の number を、getname n と入力された場合は
// n 番目に作成された社員の name を出力してください。
// 入力される値
// N
// S_1
// ...
// S_N

// ・ 1 行目では、与えられる入力の回数 N が与えられます。
// 続く N 行では、次のいずれかの形式の入力が与えられます。
// ・ make number name
// ・ getnum n
// ・ getname n

// 入力値最終行の末尾に改行が１つ入ります。
// 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
// 期待する出力
// 入力に応じた出力をしてください。
// 各入力に対する出力の末尾には改行を入れてください。

// 条件
// ・ 1 ≦ N ≦ 20 行数の範囲
// ・ 1 ≦ number ≦ 10^5  社員番号の範囲
// ・ number , name は重複しない
// ・ name は長さ 20 文字未満の文字列
// ・ 1 ≦ n ≦ (その入力時点での社員数)

// 入力例1
// 3
// 行数
// make 1 nana
// getnum 1
// getname 1

// 出力例1
// 1
// nana

// 入力例2
// 7
// make 2742 mako
// getnum 1
// make 2782 taisei
// getname 2
// make 31 megumi
// getname 1
// getname 3

// 出力例2
// 2742
// make 2742 mako getnum 1 の出力結果

// taisei
// make 2782 taisei getname 2 の出力結果

// mako
// getname 1 の出力結果 一行目の社員の名前

// megumi

// 型	%vに対応する書式指定子
// bool	%t
// int,int8 etc	%d
// uint, uint8 etc.	%d (%#vのときは%#x)
// float32, complex64, etc	%g
// string	%s
// chan	%p
// pointer	%p

import (
	"fmt"
)

// Employee クラスの定義
type Employee struct {
	number int
	name   string
}

// Employees 切り替え
var employees []Employee //! 構造体の配列をここで作成

// 社員を作成する関数
func makeEmployee(number int, name string) {
	employee := Employee{number: number, name: name} //! 構造体の各変数に入れる値を定義
	employees = append(employees, employee)          //! appendで配列に値をそれぞれ入れる
}

// n番目の社員の社員番号を取得する関数
func getNum(n int) int {
	return employees[n-1].number //! n-1はプログラミングの数の数え方上 0から数え始めるため
}

// n番目の社員の名前を取得する関数
func getName(n int) string {
	return employees[n-1].name
}

func main() {
	// 具体的な値で社員を作成
	makeEmployee(1, "John")
	makeEmployee(2, "Alice")
	makeEmployee(3, "Bob")

	// 具体的な値で社員番号と名前を取得
	fmt.Println(getNum(2))  // 出力: 2
	fmt.Println(getName(3)) // 出力: Bob
}
