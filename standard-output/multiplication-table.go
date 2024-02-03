// Go言語で九九の計算をするやり方
package main

import "fmt"

func main() {
    // 九九の表を生成して出力
    for i := 1; i <= 9; i++ {
        for j := 1; j <= 9; j++ {
            result := i * j
            fmt.Printf("%d", result)

            // 各行の最後の要素以外にはスペースを追加
            if j != 9 {
                fmt.Print(" ")
            }
        }
        fmt.Println() // 行を終了するときに改行
    }
}


