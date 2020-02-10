package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	Name string
	Age  int
}

func main() {
	f := Foo{"Emir", 22}

	// Generate json
	jsn, err := json.Marshal(f)
	if err != nil {
		fmt.Println("Couldn't create Json format")
	}

	// Print json
	fmt.Println(string(jsn))

	// Assign results to pointed f
	json.Unmarshal(jsn, &f)

	// Use structured data
	fmt.Println("Name: ", f.Name)
	fmt.Println("Age:  ", f.Age)

}
