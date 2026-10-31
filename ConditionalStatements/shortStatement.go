package main

import "fmt"

func main() {
	items := 3
	price := 49

	if total := items * price; total > 100 {
		fmt.Print("Eligible for shipping!!")
	} else {
		fmt.Println("The item is not eligible for shipping!")
	}
}
