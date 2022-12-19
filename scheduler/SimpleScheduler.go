package scheduler

import "crawl/engine"

type SimpleSchedule struct {
	workerChan chan engine.Request
}

func (s *SimpleSchedule) Submit(r engine.Request) {
	// 防止卡住
	go func() {
		s.workerChan <- r
	}()

}
func (s *SimpleSchedule) ConfigureWorkerChan(c chan engine.Request) {
	s.workerChan = c
}
