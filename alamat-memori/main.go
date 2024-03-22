package main

import "fmt"

func main() {
	var name = "Kenji"
	fmt.Println("name: ", name)
	fmt.Println("alamat dari name: ", &name) // fungsi simbol & = untuk mengetahui alamat memori

	var age = 25
	fmt.Println("age: ", age)
	fmt.Println("alamat dari age: ", &age)

}
