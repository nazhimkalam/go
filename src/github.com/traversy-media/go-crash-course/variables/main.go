package main
import "fmt"

func main() {
	var name = "Nazhim"
	var age = 19
	// var isCool = true
	isCool := true  // This is the Shorthand method for declaring a variable
	isCool = !isCool
	a,b := 1, 5

	size := 1.3

	fmt.Println(name, age+1, isCool, size)
	fmt.Println(a, b)
	fmt.Printf("%T\n", isCool)
}