package main

import (
	"awesomeProject/engine"
	"awesomeProject/parse"
	"awesomeProject/scheduler"
)

func main() {
	//const str = `<a href="https://book.douban.com/subject/6798611/" title="史蒂夫·乔布斯传"`
	//re := regexp.MustCompile(`<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`)
	//match := re.FindString(str)
	//fmt.Printf("%s", match)

	//engine.Run(engine.Request{
	//	Url: "https://book.douban.com/",
	//	ParseFunc: parse.ParseTag,
	//})

	(&engine.Concurrentengine{
		&scheduler.QueueScheduler{},
		100,
	}).Run(engine.Request{
		Url:       "https://book.douban.com/",
		ParseFunc: parse.ParseTag,
	})
}
