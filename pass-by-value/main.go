package main

import "fmt"

func main() {
	var x = 4
	var y = x // pass-by-value = menyalin nilai variabel ke variabel lainnya dan dengan memori yg berbeda

	y = y + 1
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)

	fmt.Println("alamat memori x: ", &x)
	fmt.Println("alamat memori y: ", &y)

	fmt.Println("=======================================================")

	var a = 3
	increase(a)

	fmt.Println("a:", a)
}

func increase(n int) {
	n = n + 1
	fmt.Println("n: ", n)
}
