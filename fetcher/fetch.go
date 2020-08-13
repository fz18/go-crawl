package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
	"time"
)

func BaseFetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 转码
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("error status code: %d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := DeterminEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	// 输出
	return ioutil.ReadAll(utf8Reader)
}

var ratelimit = time.Tick(100 * time.Millisecond)

func WebFetch(url string) ([]byte, error) {
	<-ratelimit
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error get url: %s", url)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36")
	resp, err := client.Do(req)

	// 转码
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("error status code: %d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := DeterminEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	// 输出
	return ioutil.ReadAll(utf8Reader)
}

func Fetch(weburl string) ([]byte, error) {
	<-ratelimit

	proxy := func(_ *http.Request) (*url2.URL, error) {
		return url2.Parse("http://127.0.0.1:1087")
	}
	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{Transport: transport}

	req, err := http.NewRequest("GET", weburl, nil)
	if err != nil {
		return nil, fmt.Errorf("ERROR: get url: %s", weburl)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36")
	resp, err := client.Do(req)

	// 转码
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("error status code: %d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := DeterminEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	// 输出
	return ioutil.ReadAll(utf8Reader)
}

func DeterminEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("fetch error : %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
