package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

func checkAndSave(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is down", url)
	} else {
		defer resp.Body.Close()
		fmt.Printf("%s Status code: %d\n", url, resp.StatusCode)
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("err", err)
		}
		file := strings.Split(url, "//")[1]
		file += ".html"

		err = ioutil.WriteFile(file, bytes, 0664)
		if err != nil {
			fmt.Println("write file error:", err)
		}
	}
	wg.Done()
}

func main() {
	urls := []string{
		"http://www.google.com",
		"http://www.golang.com",
		"http://www.youtube.com",
		"http://www.facebook.com",
		"http://www.baidu.com",
		"http://www.yahoo.com",
		"http://www.amazon.com",
		"http://www.wikipedia.org",
		"http://www.qq.com",
		"http://www.google.co.in",
		"http://www.twitter.com",
		"http://www.live.com",
		"http://www.taobao.com",
		"http://www.bing.com",
		"http://www.instagram.com",
		"http://www.weibo.com",
		"http://www.sina.com.cn",
		"http://www.linkedin.com",
		"http://onlinesbi.com",
	}
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, v := range urls {
		go check(v, &wg)
	}
	wg.Wait()
}

func check(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("%s is down.", url)
	} else {
		fmt.Printf("%s ---> stattus code: %d\n", url, resp.StatusCode)
	}
	wg.Done()

}
