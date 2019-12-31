package main

import (
	"fmt"
	"github.com/cy422396350/crowller/config"
	"github.com/cy422396350/crowller/engine"
	"github.com/cy422396350/crowller/itemServer"
	"github.com/cy422396350/crowller/scheduler"
	"github.com/cy422396350/crowller/zhenai/parser"
)

const seed = "http://www.zhenai.com/zhenghun"

func main() {
	saver, err := itemServer.CreateItemServer(fmt.Sprintf(":%d", config.Host))

	if err != nil {
		panic(err)
	}

	e := &engine.QueueEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 5,
		ItemChan:    saver,
	}

	e.Run(engine.Request{Url: seed, Parser: parser.ParserCityList})
}
