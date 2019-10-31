package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	//concurPtr := flag.Int("c", 4, "Concurrency")
	inPtr := flag.String("i", "waybackurls.txt", "WaybackUrl List")
	outPtr := flag.String("o", "waybackurls.out", "Output Filename")
	//screenShotPtr := flag.Bool("s", false, "Screenshot")
	flag.Parse()

	fileIn, err := os.Open(*inPtr)
	if err != nil {
		fmt.Println("Input file not found.")
		os.Exit(-1)
	}
	fileOut, err := os.Create(*outPtr)
	defer fileIn.Close()
	defer fileOut.Close()

	if err != nil {
		return
	}
	scannerIn := bufio.NewScanner(fileIn)
	scannerOut := bufio.NewWriter(fileOut)

	for scannerIn.Scan() {
		StatusCode, url, ContentLength := testGet(scannerIn.Text())
		fmt.Println(StatusCode, ContentLength, url)
		writeFile(StatusCode, url, ContentLength, *scannerOut)
	}
}

func writeFile(StatusCode int, url string, ContentLength int64, fileOut bufio.Writer) {
	line := strconv.Itoa(StatusCode)
	line += " " + url + " "
	line += strconv.FormatInt(ContentLength, 10)
	fileOut.WriteString(line + "\n")
	fileOut.Flush()
}

func testGet(url string) (int, string, int64) {
	response, err := http.Get(url)
	if err != nil {
		return 0, "CouldNotReq", 0
	}
	return response.StatusCode, url, response.ContentLength
}
