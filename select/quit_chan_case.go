/*
 * 一个select监听信道数据流动并对退出信号作处理的例子
 */
package main

import (
    "fmt"
)


func main() {

    c, quit := make(chan int), make(chan int)

    go func() {
        c <- 2  // 添加数据
        quit <- 1 // 发送完成信号
    } ()

    for is_quit := false; !is_quit; {
        select { // 监视信道c的数据流出
        case v := <-c: fmt.Printf("received %d from c", v)
        case <-quit: is_quit = true
        }
    }
}
