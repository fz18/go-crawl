package main

import (
	"awesomeProject/engine"
	"awesomeProject/parse/zhenai"
	"awesomeProject/persist"
	"awesomeProject/scheduler"
)

func main() {
	//const str = "qwe"
	//re := regexp.MustCompile(`<div class="m-btn purple">([\d]+)cm</div>`)
	//match := re.FindString(str)
	//fmt.Printf("%s", match)

	//engine.Run(engine.Request{
	//	Url: "https://book.douban.com/",
	//	ParseFunc: parse.ParseTag,
	//})

	(&engine.Concurrentengine{
		&scheduler.QueueScheduler{},
		100,
		persist.ItemSave(),
	}).Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: zhenai.ParseCity,
	})
}
