package main

import "fmt"

func main() {
	var name string
	name = "Emir"
	nick := "0xF61"

	fmt.Println(name, nick)
	for i := 4; i < 10; i++ {
		if i == 7 {
			fmt.Println("se7en")
		}
		fmt.Println(i)
	}
}
