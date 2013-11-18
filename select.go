package main

import (
    "fmt"
)


func foo(i int) chan int {
    c := make(chan int)
    go func () { c <- i }()
    return c
}


func main() {
    c1, c2, c3 := foo(1), foo(2), foo(3)

    c := make(chan int)

    go func() { // 开一个goroutine监视各个信道数据输出并收集数据到信道c
        for {
            select { // 监视c1, c2, c3的流出，并全部流入信道c
            case v1 := <- c1: c <- v1
            case v2 := <- c2: c <- v2
            case v3 := <- c3: c <- v3
            }
        }
    }()

    // 阻塞主线，取出信道c的数据
    for i := 0; i < 3; i++ {
        fmt.Println(<-c) // 从打印来看我们的数据输出并不是严格的1,2,3顺序
    }
}
