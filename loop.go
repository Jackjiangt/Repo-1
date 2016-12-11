package main
import (
	"fmt"
		)
loop() {
    for i := 0; i < 10; i++ {
        fmt.Printf("%d ", i)
	}
}


func main() {
	loop()
	loop()
}
