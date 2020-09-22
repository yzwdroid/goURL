package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

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

func checkStatus(link string) {
	resp, err := http.Head(link)
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.StatusCode == 200 {
		color.Green.Println(link, "is alive")
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println(`usage called
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

	for _, u := range urls {
		checkStatus(u)
	}
	// var wg sync.WaitGroup
}
