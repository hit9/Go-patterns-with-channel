package main

import "fmt"


func xrange() chan int{ // xrange用来生成自增的整数
    var ch chan int = make(chan int)

    go func() { // 开出一个goroutine
        for i := 0; ; i++ {
            ch <- i  // 直到信道索要数据，才把i添加进信道
        }
    }()

    return ch
}

func main() {

    generator := xrange()

    for i:=0; i < 1000; i++ {  // 我们生成1000个自增的整数！
        fmt.Println(<-generator)
    }
}
