package main

import "fmt"

func main() {
	fmt.Println(greeting("Nazhim"))
	fmt.Println(getSum(5, 6))
}

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2	
}