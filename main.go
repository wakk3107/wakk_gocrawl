package main

import (
	"crawl/engine"
	"crawl/parse"
	"crawl/persist"
	"crawl/scheduler"
)

func main() {
	e := &engine.ConcurrentEngine{
		10,
		&scheduler.QueueScheduler{},
		persist.ItemSave(),
	}
	e.Run(engine.Request{
		Url:       "https://book.douban.com/",
		ParseFunc: parse.ParseTag,
	})

}
