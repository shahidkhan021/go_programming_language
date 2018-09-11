package main

import "fmt"

func main() {
	// one way of decllaring maps
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	//another way of creating map
	// colors map[string]string
	//another way of creating maps
	colors := make(map[int]string)
	colors[10] = "#fffff"
	delete(colors, 10)
	fmt.Println(colors)
}
