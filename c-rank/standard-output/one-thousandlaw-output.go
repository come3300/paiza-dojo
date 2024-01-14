package main

import "fmt"

func main() {
    var numbers []int
    for i := 1; i <= 1000; i++ {
        numbers = append(numbers, i)
    }

    // fmt.Printlnを使用して各要素を改行して表示
    for _, num := range numbers {
        fmt.Println(num)
    }
}


// fmt.Printlnを使用して文字列や変数の値を標準出力に表示しています。
// 末尾に改行がされる

// fmt.Println の他の主な関数:
// fmt.Print: Print関数もPrintlnと同様に値を出力しますが、末尾に改行が追加されません。


// fmt.Print("Hello, ")
// fmt.Print("World!")
// fmt.Printf: Printfは、フォーマット文字列を使用して値を出力します。C言語のprintf関数と似ています。



// name := "Bob"
// age := 30
// fmt.Printf("Name: %s, Age: %d\n", name, age)
// %sや%dはそれぞれ文字列や整数を挿入するためのフォーマット指定子です。

// fmt.Sprintf: Sprintfはフォーマットされた文字列を生成しますが、標準出力ではなく文字列として返します。



// formattedString := fmt.Sprintf("Name: %s, Age: %d", name, age)
// fmt.Println(formattedString)