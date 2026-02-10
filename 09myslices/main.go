package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Banana", "Peach"}

	// fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Papaya")
	// fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	// fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 467
	highScores[3] = 867

	highScores = append(highScores, 777)

	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	// fmt.Println(highScores)

	// how to remove a value from slices based on index

	var courses = []string{"react", "AI", "Java", "Golang", "devops"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
