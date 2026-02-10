package main

import "fmt"

func main() {

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("PY is short for the: ", languages["PY"])

	delete(languages, "RB")
	fmt.Println("List of all languages:", languages)

	// Loops
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
