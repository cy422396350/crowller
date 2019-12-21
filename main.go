package main

import (
	"github.com/cy422396350/crowller/engine"
	"github.com/cy422396350/crowller/scheduler"
	"github.com/cy422396350/crowller/zhenai/parser"
)

const seed = "http://www.zhenai.com/zhenghun"
const shanghai = "http://www.zhenai.com/zhenghun/shanghai"

func main() {
	e := &engine.QueueEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 5,
	}
	//e.Run(engine.Request{Url:seed,Parser:parser.ParserCityList})
	e.Run(engine.Request{
		Url:    shanghai,
		Parser: parser.FindPeople,
	})
}
