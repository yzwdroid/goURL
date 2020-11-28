package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"sync"
	"time"

	"mvdan.cc/xurls/v2"

	"github.com/gookit/color"
)

type urlStatus struct {
	URL    string
	Status int
}

type post struct {
	ID  string
	URL string
}

func extractURL(str string) []string {
	rxStrict := xurls.Strict()
	foundUrls := rxStrict.FindAllString(str, -1)
	return foundUrls
}

func dataTelscope() []byte {
	var data []byte
	var posts []post
	resp, err := http.Get("http://localhost:3000/posts")
	check(err)
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		check(err)
		if err := json.Unmarshal(body, &posts); err != nil {
			panic(err)
		}
		for _, p := range posts {
			resp, err := http.Get("http://localhost:3000" + p.URL)
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				bodyData, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
				}
				data = append(data, bodyData...)
			}
		}
	}

	return data
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

func getStatusFromLink(link string) (int, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Head(link)
	if err != nil {
		color.Gray.Println(link, "is unknown")
		return 0, err
	}
	return resp.StatusCode, nil
}

func checkStatus(link string, failOnly bool, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := getStatusFromLink(link)
	if err != nil {
		fmt.Println(err)
	}

	// if in fail only mode only return the urls that returned > 400
	if failOnly && resp < 400 {
		return
	}

	switch resp {
	case 200:
		color.Green.Println(resp, link, "is alive, [OK]")
	case 300:
		color.Yellow.Println(resp, link, "it's alive, [Multiple Choices]")
	case 301:
		color.Yellow.Println(resp, link, "it's alive, [Found but its moved permanently]")
	case 307:
		color.Yellow.Println(resp, link, "it's alive, [Found but its a temporary redirect]")
	case 308:
		color.Yellow.Println(resp, link, "it's alive, [Found but its a permanent redirect]")
	case 400:
		color.Red.Println(resp, link, "is bad, [Bad Request]")
	case 401:
		color.Red.Println(resp, link, "is bad, [Unauthorized]")
	case 402:
		color.Red.Println(resp, link, "is bad, [Payment Required]")
	case 403:
		color.Red.Println(resp, link, "is bad, [Forbidden]")
	case 404:
		color.Red.Println(resp, link, "is bad, [Not Found]")
	case 410:
		color.Red.Println(resp, link, "is bad, [Gone]")
	case 500:
		color.Red.Println(resp, link, "is bad, [Internal Server Error]")
	default:
		color.Gray.Println(resp, link, "is unknown")
	}
}

func checkStatusNoColor(link string, failOnly bool, wg *sync.WaitGroup) {
	defer wg.Done()

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Head(link)
	if err != nil {
		fmt.Println(link, "is unknown")
		return
	}

	// if in fail only mode only return the urls that returned > 400
	if failOnly && resp.StatusCode < 400 {
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

func checkStatusJSON(link string, ch chan urlStatus) {

	us := urlStatus{link, 0}
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Head(link)
	if err != nil {
		ch <- us
		return
	}
	us.Status = resp.StatusCode
	ch <- us
}

func ignoreURL(f string) []string {
	var urls []string
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("^(#|https?://)")
	for scanner.Scan() {
		if !re.Match(scanner.Bytes()) {
			fmt.Println("Ignore file invalid")
			os.Exit(1)
		}
		if line := scanner.Text(); string(line[0]) != "#" {
			urls = append(urls, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return urls
}
