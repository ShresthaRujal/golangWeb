package main

import (
	"fmt"
)

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }
	colors := make(map[string]string)

	colors["white"] = "#ffffff"
	//delete
	delete(colors, "white")
	fmt.Printf("%+v", colors)
}
