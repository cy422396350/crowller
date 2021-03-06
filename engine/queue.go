package engine

import (
	"log"
)

type QueueEngine struct {
	Scheduler   Scheduler
	ItemChan    chan Item
	WorkerCount int
}
type ReadyInter interface {
	Ready(chan Request)
}

var requested = make(map[string]bool)

func (e *QueueEngine) Run(seeds ...Request) {
	out := make(chan Result)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		createQueueWorker(e.Scheduler.GetWorkerChan(), out, e.Scheduler)
	}

	for _, request := range seeds {
		if !requested[request.Url] {
			requested[request.Url] = true

			e.Scheduler.Submit(request)
		}
	}

	sum := 0
	for {
		res := <-out
		if res.Requests == nil {
			for _, item := range res.Items {
				e.ItemChan <- item
				log.Printf("get item %d\n", sum)
				sum++
			}
		}
		for _, request := range res.Requests {
			if !requested[request.Url] {
				requested[request.Url] = true
				e.Scheduler.Submit(request)
			}
		}

	}
}

/**
 * 把ready这个方法独立成一个接口,代码就简单了
因为Simple没有ready,只有Queue有
*/
func createQueueWorker(myChan chan Request, out chan Result, s ReadyInter) {
	go func() {
		for {
			s.Ready(myChan)
			request := <-myChan
			result, err := worker(request)
			if err != nil {
				log.Printf("err%v", err)
				continue
			}
			out <- result
		}
	}()
}
