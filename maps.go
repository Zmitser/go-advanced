package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red": "#ff0000",
		"black": "#000000",
		"white": "#ffffff",
	}
	//var emptyMap map[string]string
	emptyMap := make(map[string] string)
	emptyMap["key"] = "value"
	fmt.Println(colors)
	fmt.Println(emptyMap)
	delete(emptyMap, "key")
	fmt.Println(emptyMap)

	for color, hex := range colors{
		fmt.Println(color, hex)
	}
}