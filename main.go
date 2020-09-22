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
		Timeout: 2 * time.Second,
	}
	resp, err := client.Head(link)
	if err != nil {
		fmt.Println(link, "is unknown")
		return
	}
	switch resp.StatusCode {
	case 200:
		color.Green.Println(link, "is alive")
	case 400, 404:
		color.Red.Println(link, "is bad")
	default:
		fmt.Println(link, "is unknown")
	}
}

// pflag supports -v or --version
var version = flag.BoolP("version", "v", false, "print out version info")

func main() {
	flag.Parse()
	if *version == true {
		fmt.Println("goURL version 0.1")
		return
	}

	if len(os.Args) == 1 {
		fmt.Println(`
usage: ./goURL filename
example: ./goURL urls.txt, ./goURL *.txt
goURL -v or --version check version.
		`)
		os.Exit(-1)
	}

	fmt.Println(os.Args[1:])
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
	println(len(urls))
	urls = removeDuplicate(urls)
	println(len(urls))

	// wait for multiple goroutines to finish
	var wg sync.WaitGroup
	for _, u := range urls {
		wg.Add(1)
		go checkStatus(u, &wg)
	}
	wg.Wait()
}
