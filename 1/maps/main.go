package main

import "fmt"

func main() {

	/*var names map[string]int

	names = make(map[string]int, 0)

	names["Mehmet"] = 1
	names["Can"] = 2
	names["Ermurat"] = 3

	fmt.Println(names)*/

	names := map[string]int{
		"Mehmet":  1,
		"Can":     2,
		"Ermurat": 3,
	}

	delete(names, "Mehmet")

	fmt.Println(names)
}
