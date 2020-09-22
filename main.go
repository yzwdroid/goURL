package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gookit/color"
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
		Timeout: 6 * time.Second,
	}
	resp, err := client.Head(link)
	if err != nil {
		color.Red.Println(link)
		return
	}
	switch resp.StatusCode {
	case 200:
		color.Green.Println(link, "is alive")
	case 400, 404:
		color.Cyan.Println(link, "is bad")
	default:
		fmt.Println(link)
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println(`
usage called
goURL filename
goURL -v filename
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
	// var ip *int = flag.Int("flagname", 1234, "help message for flagname")
	// flag.Parse

	rxStrict := xurls.Strict()
	// urls is a slice of strings
	urls := rxStrict.FindAllString(string(dat), -1)
	println(len(urls))
	urls = removeDuplicate(urls)
	println(len(urls))
	color.Green.Println("hello world")

	var wg sync.WaitGroup
	for _, u := range urls {
		wg.Add(1)
		go checkStatus(u, &wg)
	}
	wg.Wait()
	// var wg sync.WadditGroup
}
