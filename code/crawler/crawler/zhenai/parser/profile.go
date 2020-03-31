package parser

import (
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"regexp"
	"strconv"
)

var marriageRe = regexp.MustCompile(`<div [^>]*>([离异未婚丧偶]+)</div>`)
var ageRe = regexp.MustCompile(`<div [^>]*>([\d]+)岁</div>`)
var heightRe = regexp.MustCompile(`<div [^>]*>([\d]+)cm</div>`)
var incomeRe = regexp.MustCompile(`<div [^>]*>月收入:([^<]+)</div>`)
var educationRe = regexp.MustCompile(`<div [^>]*>([大学本科高中硕士中专]+)</div>`)
var hokouRe = regexp.MustCompile(`<div [^>]*>籍贯:([^<]+)</div>`)
var xinzuoRe = regexp.MustCompile(`<div [^>]*>(.2座[^<]+)</div>`)
var houseRe = regexp.MustCompile(`<div [^>]*>工作地:([^<]+)</div>`)
var carRe = regexp.MustCompile(`<div [^>]*>([已未买车]+)</div>`)

var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func parseProfile(contents []byte, url string, name string) engine.ParseResult {

	//fmt.Printf("%s ",contents)

	exStr := extractString(contents)
	exInt := extractInt(contents)

	profile := model.Profile{
		Name:      name,
		Age:       exInt(ageRe),
		Height:    exInt(heightRe),
		Income:    exStr(incomeRe),
		Marriage:  exStr(marriageRe),
		Education: exStr(educationRe),
		Hokou:     exStr(hokouRe),
		Xinzuo:    exStr(xinzuoRe),
		House:     exStr(houseRe),
		Car:       exStr(carRe),
	}

	return engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Id:      extractString([]byte(url))(idUrlRe), // 从url中获取
				Type:    "zhenai",
				Payload: profile,
			},
		},
	}
}

func extractString(contents []byte) func(*regexp.Regexp) string {
	return func(re *regexp.Regexp) string {
		match := re.FindSubmatch(contents)
		if len(match) >= 2 {
			return string(match[1])
		}
		return ""
	}
}

func extractInt(contents []byte) func(*regexp.Regexp) int {
	return func(regexp *regexp.Regexp) int {
		matchString := extractString(contents)(regexp)
		if matchString != "" {
			age, err := strconv.Atoi(matchString)
			if err == nil {
				return age
			}
		}
		return -1
	}
}

type ProfileParser struct {
	userName string
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}

// 实现接口
func (p *ProfileParser) Parse(content []byte, url string) engine.ParseResult {
	return parseProfile(content, url, p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "ProfileParser", p.userName
}
