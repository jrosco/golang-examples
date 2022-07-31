package main

import "fmt"

func main() {
	code := map[string]int{}

	code["one"] = 1
	code["two"] = 2

	fmt.Println(code["two"])
}
