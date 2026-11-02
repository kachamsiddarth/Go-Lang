package main

import "fmt"

func main() {
	//Slices are common collection type
	//These are dynamic in nature
	//They can grow
	//[]type{....}
	names := []string{"Siddarth", "Toad", "Tony"}
	fmt.Println(names)
	names = append(names, "Racoon")
	fmt.Println(names)
}
