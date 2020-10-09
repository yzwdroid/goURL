package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gookit/color"
	flag "github.com/spf13/pflag"
	"mvdan.cc/xurls/v2"
)

type urlStatus struct {
	URL    string
	Status int
}

func removeDuplicate(urls []string) []string {
	result := make([]string, 0, len(urls))
	temp := map[string]struct{}{}
	for _, item := range urls {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func checkStatus(link string, wg *sync.WaitGroup) {
	defer wg.Done()

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Head(link)
	if err != nil {
		color.Gray.Println(link, "is unknown")
		return
	}
	switch resp.StatusCode {
	case 200:
		color.Green.Println(resp.StatusCode, link, "is alive, [OK]")
	case 300:
		color.Yellow.Println(resp.StatusCode, link, "it's alive, [Multiple Choices]")
	case 301:
		color.Yellow.Println(resp.StatusCode, link, "it's alive, [Found but its moved permanently]")
	case 307:
		color.Yellow.Println(resp.StatusCode, link, "it's alive, [Found but its a temporary redirect]")
	case 308:
		color.Yellow.Println(resp.StatusCode, link, "it's alive, [Found but its a permanent redirect]")
	case 400:
		color.Red.Println(resp.StatusCode, link, "is bad, [Bad Request]")
	case 401:
		color.Red.Println(resp.StatusCode, link, "is bad, [Unauthorized]")
	case 402:
		color.Red.Println(resp.StatusCode, link, "is bad, [Payment Required]")
	case 403:
		color.Red.Println(resp.StatusCode, link, "is bad, [Forbidden]")
	case 404:
		color.Red.Println(resp.StatusCode, link, "is bad, [Not Found]")
	case 410:
		color.Red.Println(resp.StatusCode, link, "is bad, [Gone]")
	case 500:
		color.Red.Println(resp.StatusCode, link, "is bad, [Internal Server Error]")
	default:
		color.Gray.Println(resp.StatusCode, link, "is unknown")
	}
}

func checkStatusNoColor(link string, wg *sync.WaitGroup) {
	defer wg.Done()

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Head(link)
	if err != nil {
		fmt.Println(link, "is unknown")
		return
	}
	switch resp.StatusCode {
	case 200:
		fmt.Println(resp.StatusCode, link, "is alive, [OK]")
	case 300:
		fmt.Println(resp.StatusCode, link, "it's alive, [Multiple Choices]")
	case 301:
		fmt.Println(resp.StatusCode, link, "it's alive, [Found but its moved permanently]")
	case 307:
		fmt.Println(resp.StatusCode, link, "it's alive, [Found but its a temporary redirect]")
	case 308:
		fmt.Println(resp.StatusCode, link, "it's alive, [Found but its a permanent redirect]")
	case 400:
		fmt.Println(resp.StatusCode, link, "is bad, [Bad Request]")
	case 401:
		fmt.Println(resp.StatusCode, link, "is bad, [Unauthorized]")
	case 402:
		fmt.Println(resp.StatusCode, link, "is bad, [Payment Required]")
	case 403:
		fmt.Println(resp.StatusCode, link, "is bad, [Forbidden]")
	case 404:
		fmt.Println(resp.StatusCode, link, "is bad, [Not Found]")
	case 410:
		fmt.Println(resp.StatusCode, link, "is bad, [Gone]")
	case 500:
		fmt.Println(resp.StatusCode, link, "is bad, [Internal Server Error]")
	default:
		fmt.Println(resp.StatusCode, link, "is unknown")
	}
}

// pflag supports -v or --version
var version = flag.BoolP("version", "v", false, "print out version info")

func main() {
	flag.Parse()
	if *version {
		fmt.Println("goURL version 0.1")
		return
	}

	if len(os.Args) == 1 {
		fmt.Println(`
name: goRUL
usage: go run main.go filenames
example: go run main.go urls.txt; go run main.go *.txt
go run main.go -v or --version check version.
		`)
		os.Exit(-1)
	}

	//fmt.Println("The files need to check", os.Args[1:])
	var dat []byte
	for _, file := range os.Args[1:] {
		d, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		dat = append(dat, d...)
	}

	// use xurls tool to exact links from file. Strict mod only match http://
	// and https:// schema
	rxStrict := xurls.Strict()
	// urls is a slice of strings
	urls := rxStrict.FindAllString(string(dat), -1)
	urls = removeDuplicate(urls)

	// wait for multiple goroutines to finish
	var wg sync.WaitGroup
	for _, u := range urls {
		wg.Add(1)
		if os.Getenv("CLICOLOR") == "1" {
			go checkStatus(u, &wg)
		} else if os.Getenv("CLICOLOR") == "0" {
			go checkStatusNoColor(u, &wg)
		} else {
			panic("Please set your CLICOLOR env variable.")
		}
	}
	wg.Wait()
}
