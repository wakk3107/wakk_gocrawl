package scheduler

import "crawl/engine"

type SimpleSchedule struct {
	workerChan chan engine.Request
}

func (s *SimpleSchedule) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleSchedule) WorkReady(requests chan engine.Request) {
	return
}

func (s *SimpleSchedule) Submit(r engine.Request) {
	// 防止卡住
	go func() {
		s.workerChan <- r
	}()

}
func (s *SimpleSchedule) WorkChan() chan engine.Request {
	return s.workerChan
}
