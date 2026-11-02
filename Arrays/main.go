package main

import "fmt"

func main() {
	var marks [3]int
	marks[0] = 10
	marks[1] = 30
	marks[2] = 50

	fmt.Println(marks)
	fmt.Println(len(marks))
	//ARRAY LITERAL
	res := [5]int{2, 3, 4, 5, 6}
	fmt.Println(res)

}
