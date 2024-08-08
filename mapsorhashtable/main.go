package main

import (
	"fmt"
)

func main() {

	//we can use make to create slices and maps
	languages := make(map[string]string)
	languages["JS"] = "javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	fmt.Println("list of all languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])
	delete(languages, "RB")
	fmt.Println("list of all languages: ", languages)

	//How to loop throught the maps
	//loops are interesting in Go lang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
