package main
import "fmt"

func main() {
	names := [3]string{"Nazhim", "Hashim", "Kalam"}
	fmt.Println(names[2])
	fmt.Println(len(names))
	fmt.Println(names[1:3])
}