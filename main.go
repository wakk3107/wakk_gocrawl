package main

import (
	"crawl/engine"
	"crawl/parse"
)

func main() {
	engine.Run(engine.Request{
		Url:       "https://book.douban.com/tag/%E5%B0%8F%E8%AF%B4",
		ParseFunc: parse.ParseBookList,
	})
}
