package main

import "fmt"

func main() {

	//var numbers = []int{1, 2, 3, 4}

	/*for index := 0; index < len(numbers); index++ {
		fmt.Println(numbers[index])
	}*/

	/*for index, value := range numbers {
		fmt.Println(index, value)
	}*/

	/*var language = "Golang"

	for _, character := range language {
		fmt.Printf("Character: %c\n", character)
	}*/

	var names = map[string]int{
		"Mehmet":  10,
		"Can":     20,
		"Ermurat": 30,
	}

	for key, value := range names {
		fmt.Println(key, value)
	}

}
