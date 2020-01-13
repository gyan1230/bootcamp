package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func checkURL(url string, c chan string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("url is down", url)
	} else {
		fmt.Printf("%s is up and status code is %d\n", url, res.StatusCode)
		c <- url
	}
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
	c := make(chan string)
	for _, v := range urls {
		go checkURL(v, c)
	}

	// for {
	// 	go checkURL(<-c, c)
	// 	fmt.Println(strings.Repeat("*", 29))
	// 	time.Sleep(2 * time.Second)
	// }

	//same can br achieve by range over channel

	for url := range c {
		time.Sleep(2 * time.Second)
		go checkURL(url, c)
		fmt.Println(strings.Repeat("*", 29))
	}

	//same can br achieve by anonymous function over channel

	// for url := range c {
	// 	go func(u string) {
	// 		checkURL(u, c)
	// 		time.Sleep(3 * time.Second)

	// 	}(url)
	// }
}
