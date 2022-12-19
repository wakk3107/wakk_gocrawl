package main

import (
	"crawl/engine"
	"crawl/parse"
	"crawl/scheduler"
)

func main() {
	e := &engine.ConcurrentEngine{
		10,
		&scheduler.QueueScheduler{},
	}
	e.Run(engine.Request{
		Url:       "https://book.douban.com/",
		ParseFunc: parse.ParseTag,
	})

}
