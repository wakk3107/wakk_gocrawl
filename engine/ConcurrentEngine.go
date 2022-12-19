package engine

import (
	"crawl/fetcher"
	"fmt"
	"log"
)

type Scheduler interface {
	Submit(Request)
	Run()
	WorkReady(chan Request)
	WorkChan() chan Request
}
type ConcurrentEngine struct {
	WorkCount int
	Scheduler Scheduler
	ItemChan  chan interface{}
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()
	// 生成工人
	for i := 0; i < e.WorkCount; i++ {
		CreateWork(e.Scheduler.WorkChan(), out, e.Scheduler)
	}
	// 上传第一批种子
	for _, seed := range seeds {
		e.Scheduler.Submit(seed)
	}
	// 一直去接收结果
	for {
		result := <-out
		for _, item := range result.Items {
			var res = item
			go func() {
				e.ItemChan <- res
			}()
		}
		// 提交衍生任务
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}
func CreateWork(in chan Request, out chan ParseResult, s Scheduler) {

	go func() {
		for {
			// 告诉调度器当前 worker 准备就绪，任务放进 in 吧
			s.WorkReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				fmt.Println("error: ", err)
				continue
			}
			out <- result

		}
	}()
}

func worker(r Request) (ParseResult, error) {
	fmt.Println("fetch url : ", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetch Url %s error :%s", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParseFunc(body), nil

}
