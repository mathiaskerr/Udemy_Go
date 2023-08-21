package main

import "fmt"

func main(){

	// colours := make(map[string]string)

    // var colours map[string]string

	colours := map[string]string{
		"red":"#ff0000",
		"green":"#4bf745",
	}
	// adding key value pairs
	colours["white"] = "#ffffff"

	// delete(colours, "white")

	fmt.Println(colours)
	printMap(colours)
} 

func printMap(c map[string]string){
	for colour, hex := range c{
		fmt.Println("Hex code for ", colour, " is ", hex)
	}
}