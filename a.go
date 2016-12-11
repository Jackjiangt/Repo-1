package main
import (
	"fmt"
	)
func main(){
	func getmsg(i int) (r string){
		fmt.Println(i)
		r= "hi"
		return r
	}
}
