package main
//fmt is a "Format package for Go"
import (
	"fmt"
)
//creating the function
func add(a int,b int)int{ 
      return a+b
}
//addition using the function
func main() {
	fmt.Println(add(5,10))
}
