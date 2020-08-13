package zhenai

import (
	"awesomeProject/engine"
	"regexp"
)

var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[\d]+)" target="_blank">([^<]+)</a>`)
var cityUrlRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^"]+)"`)

func ParseCity(contents []byte) engine.ParseResult {
	matchs := cityRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}

	for _, m := range matchs {
		name := string(m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})
	}

	matchs = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matchs {
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseCity,
		})
	}
	return result
}
