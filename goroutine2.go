package main

import "fmt"
import "runtime"

var quit chan int = make(chan int)

func loop() {
    for i := 0; i < 100; i++ { //为了观察，跑多些
        fmt.Printf("%d ", i)
    }
    quit <- 0
}

func main() {
//    runtime.GOMAXPROCS(2) // 最多使用2个核
    runtime.GOMAXPROCS(runtime.NumCPU()) 
	fmt.Println("the number of Mac CPU Processos is", runtime.NumCPU())
    go loop()
    go loop()

    for i := 0; i < 2; i++ {
        <- quit
    }
}
