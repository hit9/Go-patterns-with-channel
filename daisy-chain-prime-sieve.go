/*
 *  利用信道菊花链筛法求某一个整数范围的素数
 *  筛法求素数的基本思想是：把从1开始的、某一范围内的正整数从小到大顺序排列，
 *  1不是素数，首先把它筛掉。剩下的数中选择最小的数是素数，然后去掉它的倍数。
 *  依次类推，直到筛子为空时结束
 */
package main

import "fmt"

func xrange() chan int{ // 从2开始自增的整数生成器
    var ch chan int = make(chan int)

    go func() { // 开出一个goroutine
        for i := 2; ; i++ {
            ch <- i  // 直到信道索要数据，才把i添加进信道
        }
    }()

    return ch
}


func filter(in chan int, number int) chan int {
    // 输入一个整数队列，筛出是number倍数的, 不是number的倍数的放入输出队列
    // in:  输入队列
    out := make(chan int)

    go func() {
        for {
            i := <- in // 从输入中取一个

            if i % number != 0 {
                out <- i // 放入输出信道
            }
        }
    }()

    return out
}


func main() {
    const max = 100 // 找出100以内的所有素数
    nums := xrange() // 初始化一个整数生成器
    number := <-nums  // 从生成器中抓一个整数(2), 作为初始化整数
    out := make(chan int) // 输出队列临时变量

    for number <= max { // number作为筛子，当筛子超过max的时候结束筛选
        fmt.Println(number) // 打印素数
        out = filter(nums, number) //筛掉number的倍数
        number = <- nums  // 继续取下一个作为筛子
        nums = out // 更新输入信道为筛选后的
    }
}
