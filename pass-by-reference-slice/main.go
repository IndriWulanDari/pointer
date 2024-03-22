package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
// 	var numbers = []int{4, 3, 7, 11}
// 	var anotherNumbers = numbers

// 	fmt.Println("numbers: ", numbers)
// 	fmt.Println("anotherNumbers: ", anotherNumbers)

// 	numbers[1] = 9
// 	fmt.Println("numbers: ", numbers)
// 	fmt.Println("anotherNumbers: ", anotherNumbers)

// 	fmt.Println("=====================================")
// 	var scores = []int{7, 8, 6, 9}
// 	multiplyBy10(scores) // pass by reference

// 	fmt.Println("scores:", scores)
// }

// // function pada argument
// func multiplyBy10(numbers []int) {
// 	for i := range numbers {
// 		numbers[i] = numbers[i] * 10
// 	}
// 	fmt.Println("numbers di function:", numbers)
// }

//Soal Quiz

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	giver, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	receiver, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	marble, _ := strconv.Atoi(scanner.Text())

	giveMarble(&giver, &receiver, marble)
	fmt.Printf("%d \n", giver)
	fmt.Printf("%d", receiver)
}

func giveMarble(giver, receiver *int, amount int) {
	if amount > *giver {
		fmt.Println("Kelereng tidak cukup untuk dibagikan")
	} else {
		*giver -= amount
		*receiver += amount
	}
}
