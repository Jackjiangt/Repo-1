package main

import "fmt"
import "runtime"

var quit chan int = make(chan int)

func loop() {
    for i := 0; i < 100; i++ { //为了观察，跑多些
		runtime.Gosched()
        fmt.Printf("%d ", i)
    }
    quit <- 0
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU()) // 最多使用2个核

    go loop()
    go loop()

    for i := 0; i < 2; i++ {
        <- quit
    }
}
