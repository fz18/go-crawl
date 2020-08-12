package parse

import (
	"awesomeProject/engine"
	"awesomeProject/model"
	"regexp"
	"strconv"
)

var autoRe = regexp.MustCompile(`<span class="pl"> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`)
var publicerRe = regexp.MustCompile(`<span class="pl">出版社:</span>([^<]+)<br/>`)
var bookpagesRe = regexp.MustCompile(`<span class="pl">页数:</span> ([^<]+)<br/>`)
var priceRe = regexp.MustCompile(`<span class="pl">定价:</span>([^<]+)<br/>`)
var scoreRe = regexp.MustCompile(`<strong class="ll rating_num " property="v:average">([^<]+)</strong>`)
var infoRe = regexp.MustCompile(`<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`)

func ParseBookDetail(contents []byte, bookName string) engine.ParseResult {

	bookDetail := model.BookDetail{}

	bookDetail.BookName = bookName
	bookDetail.Author = ExtractString(contents, autoRe)
	bookDetail.Publicer = ExtractString(contents, publicerRe)
	page, err := strconv.Atoi(ExtractString(contents, bookpagesRe))
	if err == nil {
		bookDetail.Bookpages = page
	}
	bookDetail.Info = ExtractString(contents, infoRe)
	bookDetail.Price = ExtractString(contents, priceRe)
	bookDetail.Score = ExtractString(contents, scoreRe)

	result := engine.ParseResult{
		Items: []interface{}{bookDetail},
	}
	return result
}

func ExtractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
