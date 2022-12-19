package scheduler

import "crawl/engine"

type QueueScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q *QueueScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

func (q *QueueScheduler) ConfigureWorkerChan(r chan engine.Request) {
	//TODO implement me
	panic("implement me")
}
func (q *QueueScheduler) WorkReady(w chan engine.Request) {
	q.workerChan <- w
}
func (q *QueueScheduler) Run() {
	q.workerChan = make(chan chan engine.Request)
	q.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWork chan engine.Request
			// 有任务了且有工人能去执行就取出来
			if len(requestQ) > 0 && len(workQ) > 0 {
				activeRequest = requestQ[0]
				activeWork = workQ[0]
			}
			select {
			// 更新任务队列
			case r := <-q.requestChan:
				requestQ = append(requestQ, r)
			// 更新 工人 队列
			case w := <-q.workerChan:
				workQ = append(workQ, w)
			// 派发任务
			case activeWork <- activeRequest:
				workQ = workQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
