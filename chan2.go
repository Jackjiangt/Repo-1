  package main
  import "fmt"

  func main(){
  c := make(chan int,2)
  c<- 1
  c<- 2
  v1 := <-c
  v2 := <-c
  //fmt.Println(<-c)
  //fmt.Println(<-c)
  fmt.Println(v1)
  fmt.Println(v2)
  }
