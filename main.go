package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"

	flag "github.com/spf13/pflag"
	"mvdan.cc/xurls/v2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// pflag supports -v or --version
var version = flag.BoolP("version", "v", false, "print out version info")
var js = flag.BoolP("json", "j", false, "output json format to stdout")
var fp = flag.StringP("file", "f", "", "file name to check")
var ignore = flag.BoolP("ignore", "i", false, "ignore url patterns")

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

	dat, err := ioutil.ReadFile(*fp)
	check(err)

	// use xurls tool to exact links from file. Strict mod only match http://
	// and https:// schema
	rxStrict := xurls.Strict()
	// urls is a slice of strings
	urls := rxStrict.FindAllString(string(dat), -1)
	urls = removeDuplicate(urls)

	if *ignore {
		var tp []string
		urlsAfterIgnore := ignoreURL("ignore.txt")
		for _, link := range urls {
			valid := true
			for _, url := range urlsAfterIgnore {
				if strings.HasPrefix(link, url) {
					valid = false
					break
				}
			}
			if valid {
				tp = append(tp, link)
			}
		}
		urls = tp
	}

	// wait for multiple goroutines to finish
	var wg sync.WaitGroup

	if *js {
		ch := make(chan urlStatus)
		s := make([]urlStatus, 0)

		for _, u := range urls {
			go checkStatusJSON(u, ch)
		}

		for range urls {
			it := <-ch
			s = append(s, it)
		}

		data, err := json.Marshal(s)
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		os.Stdout.WriteString(string(data))
	} else {
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
}
