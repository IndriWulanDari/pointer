package main

import "fmt"

//pass by reference pada pointer = merubah nilai pada masing2 variabel yg memiliki alamat memori yang sama

func main() {
	var x int = 4
	var y *int = &x //pass by reference

	fmt.Println("x:", x)
	fmt.Println("alamat x:", &x)
	fmt.Println("y:", y)
	fmt.Println("alamat y:", &y)

	fmt.Println("Nilai reference dari pointer y:", *y) //simbol * = pointer

	*y = *y + 1

	fmt.Println("x:", x) //melihat apakah nilai x berubah, karena yg diganti yaitu variabel y
	fmt.Println("*y:", *y)

	fmt.Println("==============================================")

	var a int = 3
	increase(&a) // pass by reference
	fmt.Println("a:", a)
}

func increase(n *int) {
	*n = *n + 1
	fmt.Println("n di function increase:", *n)
}
