package engine

import (
	"crawl/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	requests = append(requests, seeds...)
	for len(requests) > 0 {
		// 取第一个来干
		r := requests[0]
		requests = requests[1:]
		// 获取内容
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Println("error: ", err)
			return
		}
		parseResult := r.ParseFunc(body)
		// 再次更新任务栏
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got Items: %s", item)

		}

	}
}
