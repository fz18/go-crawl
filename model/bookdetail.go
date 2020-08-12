package model

import "strconv"

type BookDetail struct {
	BookName  string
	Author    string
	Publicer  string
	Bookpages int
	Price     string
	Score     string
	Info      string
}

func (b BookDetail) String() string {
	return "书名：" + b.BookName + " 作者：" + b.Author + " 出版社：" + b.Publicer + " 页数：" + strconv.Itoa(b.Bookpages) + " 价格：" + b.Price + " 分数：" + b.Score + "\n简介：" + b.Info
}
