package main

import (
	"crawl/engine"
	"crawl/parse"
)

func main() {
	engine.Run(engine.Request{
		Url:       "https://book.douban.com/subject/36104107/",
		ParseFunc: parse.ParseBookDetail,
	})
}
