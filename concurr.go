package main

import {
    "fmt"
    "runtime"
    "time"
}
var MULTICORE int
func main() {
    MULTICORE = runtime.NumCPU()    //计算出本地的cpu核总数
    //指定MULTICORE个核来运行
    //这里没有设置cpu亲和性，所以各个线程会在任意cpu核上跑，同一个线程也可能会不断跳到不同核上运行
    runtime.GOMAXPROCS(MULTICORE) 
    // 启动MULTICORE个goroutine来执行test()
    for i := 0; i < MULTICORE; i++ {
        go test()
    }
    // sleep 10s是为了让主进程等待所有goroutine都运行退出
    time.Sleep(10*time.Second)
}
func test() {
    for i := 0; i < 10; i++ {
        fmt.Printf("test\n")
    }
}
