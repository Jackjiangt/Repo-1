package mian
import (
	"fmt"
		)

type Person struct{
	name string
	age int
}

func (this *Person) SetName(name string){
	this.name = name
}

func (this *Person) GetName() string{
	return this.name
}

func (this *Person) SetAge(age int){
	this.age = age
}

func (this *Person) GetAge() int{
	return this.age
}


type Girl struct {
	Person
	gender string
}

func (this *Girl) SetGender(gender string) {
    this.gender = gender
}

func (this *Girl) GetGender() string {
	      return this.gender
		  }

func main(){
	p1 := new(Person)
	p1.SetName("Jack Jaing")
	p1.SetAge(35)
	fmt.Println("Name: %s, Age: %d\n", p1.GetName(), p1.GetAge())

	var g Girl
	g.SetName("lili")
    g.SetAge(12)
	g.SetGender("girl")
	fmt.Println("Name: %s, Age: %d, Gender: %s\n", g.GetName(), g.GetAge(), g.GetGender())

}
