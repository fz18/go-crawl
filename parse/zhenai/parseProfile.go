package zhenai

import (
	"awesomeProject/engine"
	"awesomeProject/model"
	"awesomeProject/parse"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)岁</div>`)
var marry = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">(已婚)</div>`)
var constellation = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>(天秤座)</div>`)
var height = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)cm</div>`)
var weight = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)kg</div>`)
var salary = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>月收入:([^<]+)</div>`)

var idRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func ParseProfile(contents []byte, name string) engine.ParseResult {

	profile := model.Profile{}
	profile.Name = name

	age, err := strconv.Atoi(parse.ExtractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(parse.ExtractString(contents, height))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(parse.ExtractString(contents, weight))
	if err == nil {
		profile.Weight = weight
	}

	profile.Salary = parse.ExtractString(contents, salary)
	profile.Constellation = parse.ExtractString(contents, constellation)
	if parse.ExtractString(contents, marry) == "" {
		profile.Marry = "未婚"
	} else {
		profile.Marry = "已婚"
	}

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}
