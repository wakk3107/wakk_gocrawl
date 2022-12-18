package main

import (
	"crawl/engine"
	"crawl/parse"
)

func main() {
	engine.Run(engine.Request{
		Url:       "https://book.douban.com/",
		ParseFunc: parse.ParseContent,
	})
}
