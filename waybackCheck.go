package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	concurPtr := flag.Int("c", 4, "Concurrency")
	inPtr := flag.String("i", "waybackurls.txt", "WaybackUrl List")
	outPtr := flag.String("o", "waybackurls.out", "Output Filename")
	timeoutPtr := flag.Int("t", 2, "Timeout in second")

	//screenShotPtr := flag.Bool("s", false, "Screenshot")
	flag.Parse()
	maxThreads := *concurPtr
	var wg sync.WaitGroup

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
		wg.Add(1)
		go worker(scannerIn.Text(), *timeoutPtr, *scannerOut, &wg)
		for i := 0; i < maxThreads-1; i++ {
			if scannerIn.Scan() {
				wg.Add(1)
				go worker(scannerIn.Text(), *timeoutPtr, *scannerOut, &wg)
			} else {
				break
			}
		}
		wg.Wait()
	}
}

func worker(url string, timeout int, out bufio.Writer, wg *sync.WaitGroup) {
	StatusCode, url, ContentLength := testGet(url, timeout)
	fmt.Println(StatusCode, ContentLength, url)
	writeFile(StatusCode, url, ContentLength, out, wg)
}

func writeFile(StatusCode int, url string, ContentLength int64, fileOut bufio.Writer, wg *sync.WaitGroup) {
	defer wg.Done()
	line := strconv.Itoa(StatusCode)
	line += " " + url + " "
	line += strconv.FormatInt(ContentLength, 10)
	fileOut.WriteString(line + "\n")
	fileOut.Flush()
}

func testGet(url string, timeout int) (int, string, int64) {
	var netClient = &http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}
	response, err := netClient.Get(url)
	if err != nil {
		return 0, "CouldNotReq", 0
	}
	return response.StatusCode, url, response.ContentLength
}
