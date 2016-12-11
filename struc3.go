package main
import "fmt"

type Person struct{
	name string
	age int
}

func main(){
	var p1 Person
	p2 := Person{}
	p3 := Person{"Jack Jiang",35}
	p4 := Person{age:35}
	p5 := new(Person)
	fmt.Println(p1,p2,p3,p4,p5)
	
	p1.age=25
	p2.age=30
	if compareAge(p1,p2){
		fmt.Println("p1 is older than p2")
	} else {
		fmt.Println("p2 is older than p1", p2.age)
	}

}

func compareAge(p1,p2 Person) bool{
	if (p1.age > p2.age) {
		return true
	} else {
		return false
	} 

}
