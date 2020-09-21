package main

import (
	"io/ioutil"

	"mvdan.cc/xurls/v2"
)

func main() {
	dat, err := ioutil.ReadFile("urls.txt")
	// var ip *int = flag.Int("flagname", 1234, "help message for flagname")
	// flag.Parse
	if err != nil {
		panic(err)
	}
	rxStrict := xurls.Strict()
	// urls is a slice of strings
	urls := rxStrict.FindAllString(string(dat), -1)
	for _, u := range urls {
		println(u)
	}
	println(len(urls))
	//seen := make(map[string]bool)
}
