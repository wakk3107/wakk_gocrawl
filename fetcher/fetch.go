package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// 原生的方式
func Fetch(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bodyReader := bufio.NewReader(resp.Body)
	// 探测编码
	ec := DetectEncoding(bodyReader)
	// 生成对应编码的转换阅读器
	utf8Reader := transform.NewReader(bodyReader, ec.NewDecoder())
	// 阅读器提供内容
	return ioutil.ReadAll(utf8Reader)

}

// 编码探测函数，返回探测出的编码
func DetectEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		fmt.Println("error: ", err)
		return unicode.UTF8
	}
	// 这个包也要安装
	// go get golang.org/x/net/html/charset
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e

}
