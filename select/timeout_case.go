/*
 * 一个对select监听信道作超时处理的例子
 */

package main

import (
    "fmt"
    "time"
)


func foo(i int) chan int {
    c := make(chan int)
    go func () { c <- i }()
    return c
}


func main() {
    c1, c2, c3 := foo(1), foo(2), foo(3)

    timeout := time.After(1 * time.Second) // timeout 是一个计时信道, 如果达到时间了，就会发一个信号出来

    for is_timeout := false; !is_timeout; {
        select { // 监视信道c1, c2, c3, timeout信道的数据流出
        case v1 := <- c1: fmt.Printf("received %d from c1", v1)
        case v2 := <- c2: fmt.Printf("received %d from c2", v2)
        case v3 := <- c3: fmt.Printf("received %d from c3", v3)
        case <- timeout: is_timeout = true // 超时
        }
    }
}
