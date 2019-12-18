package main

import (
	"github.com/cy422396350/crowller/engine"
	"github.com/cy422396350/crowller/zhenai/parser"
)

const seed = "http://www.zhenai.com/zhenghun"

func main() {
	engine.Run(engine.Request{Url:seed,Parser:parser.ParserCityList})
}

