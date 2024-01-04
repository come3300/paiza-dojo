package main

// クラスの学級委員である paiza 君は、クラスのみんなに次のような形式でアカウントの情報を送ってもらうよう依頼しました。

// 名前 年齢 誕生日 出身地

// 送ってもらったデータを使いやすいように整理したいと思った paiza 君はクラス全員分のデータを次のような構造体でまとめることにしました。

// student{
//     name : 名前 20文字以内
//     old : 年齢
//     birth : 誕生日 12/30形式
//     state : 出身地 20文字以内
// }

// 年齢ごとの生徒の名簿を作る仕事を任された paiza 君はクラスメイトのうち、決まった年齢の生徒を取り出したいと考えました。
// 取り出したい生徒の年齢が与えられるので、その年齢の生徒の名前を出力してください。

// 期待する出力
// 年齢が K である生徒の名前を 1 行で出力してください。

// 10人（10行）

// 名前 n_i
// 年齢 o_i    1 ≦ o_i , K ≦ 100 年齢が K であるような生徒はクラスに 1 人
// 誕生日 b_i
// 出身地  s_i 



import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Student struct {
	Name  string
	Old   int
	Birth string
	State string
}

// func main() {
// 	// 標準入力からの読み込み
// 	scanner := bufio.NewScanner(os.Stdin)

// 	// クラスメイトの情報を格納するスライス
// 	var classmates []Student

// 	// クラスメイトの情報を取得
// 	for i := 0; i < 3; i++ {
// 		scanner.Scan()
// 		line := scanner.Text()
// 		fields := strings.Fields(line)
// 		old := parseInt(fields[1])

// 		classmates = append(classmates, Student{
// 			Name:  fields[0],
// 			Old:   old,
// 			Birth: fields[2],
// 			State: fields[3],
// 		})
// 	}

// 	// 検索する年齢 K を取得
// 	scanner.Scan()
// 	k := parseInt(scanner.Text())

// 	// 年齢が K である生徒の名前を出力
// 	for _, student := range classmates {
// 		if student.Old == k {
// 			printStudentName(student)
// 			break
// 		}
// 	}
// }

// // 文字列を整数に変換するヘルパー関数
// func parseInt(s string) int {
// 	var result int
// 	fmt.Sscanf(s, "%d", &result)
// 	return result
// }

// // 生徒の名前を出力する関数
// func printStudentName(student Student) {
// 	fmt.Printf("名前: %s\n", student.Name)
// }


// go run b-rank/class-struct/step1.goで実行可能
