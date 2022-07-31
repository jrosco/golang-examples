package main

import "fmt"

func main() {
	fmt.Println("\nFor Loop:\n")

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("\n Still a for loop:\n")

	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}

	fmt.Println("\nMore while-ish:\n")

	j := 10
	for {
		if j <= 0 {
			break
		} else {
			fmt.Println("GO!", j)
		}
		j--
	}

	fmt.Println("\nThrough an iterator:\n")

	numbers := [4]string{"Zero", "One", "Two", "Three"}

	for index, number := range numbers {
		fmt.Println(index, number)
	}

	fmt.Println("\nFinished!")

}
