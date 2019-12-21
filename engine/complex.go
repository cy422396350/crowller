package engine

import (
	"log"
	"sync"
)

type ComplexEngine struct {
	Scheduler   Scheduler
	wg          *sync.WaitGroup
	WorkerCount int
}

type Scheduler interface {
	ReadyInter
	GetWorkerChan() chan Request
	Submit(Request)
	Run()
}

/**
 * scheduler实现一个调度器,getWorkerChan 获取request的channel,
 * submit 开启一个协程来防止循环阻塞
 * ReadyInter 是一个准备好的接口
 */
func (e *ComplexEngine) Run(seeds ...Request) {
	e.wg = &sync.WaitGroup{}
	out := make(chan Result)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		e.wg.Add(1)
		createWorker(e.Scheduler.GetWorkerChan(), out, e.wg)
	}
	for _, request := range seeds {
		e.wg.Add(1)
		e.Scheduler.Submit(request)
	}
	sum := 0
	go func() {
		for {
			res := <-out
			for _, item := range res.Items {
				log.Printf("GOT %d item %v", sum, item)
				sum++
			}
			for _, request := range res.Requests {
				if request.Url != "" {
					e.wg.Add(1)
				}
				e.Scheduler.Submit(request)
			}
		}

	}()
	e.wg.Wait()
}

func createWorker(in chan Request, out chan Result, wg *sync.WaitGroup) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
			wg.Done()
		}
	}()
}
