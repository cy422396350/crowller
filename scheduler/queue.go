package scheduler

import (
	"github.com/cy422396350/crowller/engine"
)

type QueueScheduler struct {
	Request  chan engine.Request
	WorkChan chan chan engine.Request
}

func (q *QueueScheduler) GetWorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueueScheduler) Ready(w chan engine.Request) {
	q.WorkChan <- w
}

func (q *QueueScheduler) Submit(r engine.Request) {
	q.Request <- r
}

/**
 * Queue是把Run进行select维持一个队列,然后不断地死循环维护队列
 * WorkerChan是一个chan chan Request 意思是一个channel里面的一个类型是Channel
 */
func (q *QueueScheduler) Run() {
	q.WorkChan = make(chan chan engine.Request)
	q.Request = make(chan engine.Request)
	go func() {
		var requests []engine.Request
		var workers []chan engine.Request
		for {
			var request engine.Request
			var worker chan engine.Request
			if len(requests) > 0 && len(workers) > 0 {
				request = requests[0]
				worker = workers[0]
			}
			select {
			case requestQ := <-q.Request:
				requests = append(requests, requestQ)
			case workerQ := <-q.WorkChan:
				workers = append(workers, workerQ)
			case worker <- request:
				requests = requests[1:]
				workers = workers[1:]
			}
		}
	}()
}
