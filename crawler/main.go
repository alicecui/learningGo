package main

import (
	"u2pppw/crawler/engine"
	"u2pppw/crawler/scheduler"
	"u2pppw/crawler/zhenai/parser"
)

func main() {
	e:= engine.ConcurrentEngine{
		Scheduler:&scheduler.SimpleScheduler{},
		WorkerCount:10,
	}

	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCityList,
	})
}
