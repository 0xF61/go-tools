package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	isEncode := flag.Bool("d", false, "Encode")
	flag.Parse()
	text, _ := reader.ReadString('\n')
	if *isEncode {
		result, err := url.QueryUnescape(text)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Print(result)
	} else {
		result := url.QueryEscape(text)
		fmt.Print(result)
	}
}
