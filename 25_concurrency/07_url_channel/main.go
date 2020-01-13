package main

import "net/http"

import "fmt"

func check(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		msg := fmt.Sprintf("%s is down\n", url)
		msg += fmt.Sprintf("%v error\n", err)
		ch <- msg
	} else {
		msg := fmt.Sprintf("%s is up and status code is %d\n", url, resp.StatusCode)
		ch <- msg
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
	ch := make(chan string)
	for _, url := range urls {
		go check(url, ch)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-ch)
	}
}
