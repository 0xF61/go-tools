package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Base struct {
	Field1 string
	Field2 string
}

func main() {
	res, err := http.Get("http://localhost:1337/ping")
	if err != nil {
		log.Fatalln(err)
	}

	var base Base
	if err := json.NewDecoder(res.Body).Decode(&base); err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	log.Printf("%s -> %s\n", base.Field1, base.Field2)
}
