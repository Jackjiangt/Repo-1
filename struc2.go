package main  
  
import "fmt"  
  
//1. 声明一个自定义类型名为 Person 的结构体  
type Person struct {  
    name string  
    age int  
}  
  
func main() {  
    //2. 初始化  
    var p1 Person  
    p2 := Person{}
    p3 := Person{"James", 23}  
    p4 := Person{age:23}  
    fmt.Println(p1, p2, p3, p4)  
    p5 := new(Person)  
    p6 := &Person{}  
    p7 := &Person{"James", 23}  
    p8 := &Person{age:23}  
    fmt.Println(p5, p6, p7, p8)  
    /*********************************/  
    /*print result                   */  
    /*{ 0} { 0} {James 23} { 23}     */  
    /*&{ 0} &{ 0} &{James 23} &{ 23} */  
    /*********************************/  
      
    //3. 操作  
    p1.age = 50  
    p2.age = 25  
    if compareAge(p1, p2) {  
        fmt.Println("p1 is older than p2")  
    } else {  
        fmt.Println("p2 is older than p1")  
    }  
    /*********************************/  
    /*print result                   */  
    /*p1 is older than p2            */  
    /*********************************/  
}  
  
func compareAge(p1, p2 Person) bool {  
    if p1.age > p2.age {  
        return true  
    }  
    return false  
}
