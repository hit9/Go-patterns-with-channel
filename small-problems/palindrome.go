/*
 * 用信道作回文字符串判断
 */

package main

import "fmt"


func is_palindrome (str string) bool { // 判断是否是回文数

    ch := make(chan byte)

    length := len(str)

    go func() {
        for i := 0; i < length; i++ {
            ch <- str[i]
        }
    }()

    for i := length-1; i >=0 ; i-- {
        if <- ch != str[i] {
            return false
        }
    }

    return true
}


func main() {
    fmt.Println(is_palindrome("hello")) // false
    fmt.Println(is_palindrome("hellooll")) // false
    fmt.Println(is_palindrome("helloolleh")) // true
}
