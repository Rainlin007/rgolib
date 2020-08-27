package main

import "fmt"

func init() {
	fmt.Println("hello init")
}
func PrintHello() {
	fmt.Println("hello")
}
func main() {
	PrintHello()
}
