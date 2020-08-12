package parse

import (
	"awesomeProject/engine"
	"regexp"
)

const BooklistRe = `<a href="([^"]+)" title="([^"]+)"`

func ParseBookList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(BooklistRe)

	matchs := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matchs {
		bookName := string(m[2])
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseBookDetail(c, bookName)
			},
		})
	}

	return result
}
