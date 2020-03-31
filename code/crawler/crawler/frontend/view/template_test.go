package view

import (
	"html/template"
	"learngo/crawler/engine"
	"learngo/crawler/frontend/model"
	common "learngo/crawler/model"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {

	// 必须合法，否则panic
	template := template.Must(
		template.ParseFiles("template.html"))

	page := model.SearchResult{
		Hits: 123,
	}

	item := engine.Item{
		Url:  "http://www.baidu.com",
		Type: "zhenai",
		Id:   "100082",
		Payload: common.Profile{
			Age:        11,
			Height:     111,
			Weight:     22,
			Income:     "333-33332",
			Gender:     "女",
			Name:       "安静的",
			Xinzuo:     "射手座",
			Occupation: "人事/行政",
			Marriage:   "未婚",
			House:      "已购房",
			Hokou:      "山东",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	// os.Stdout向控制台输出
	err := template.Execute(os.Stdout, page)
	if err != nil {
		panic(err)
	}

	// 输出为文件
	out, err := os.Create("template.test.html")
	err = template.Execute(out, page)
	if err != nil {
		panic(err)
	}
}

func TestSearchResultRender(t *testing.T) {

	view := CreateSearchResultView("template.html")

	page := model.SearchResult{
		Hits: 123,
	}

	item := engine.Item{
		Url:  "http://www.baidu.com",
		Type: "zhenai",
		Id:   "100082",
		Payload: common.Profile{
			Age:        11,
			Height:     111,
			Weight:     22,
			Income:     "333-33332",
			Gender:     "女",
			Name:       "安静的",
			Xinzuo:     "射手座",
			Occupation: "人事/行政",
			Marriage:   "未婚",
			House:      "已购房",
			Hokou:      "山东",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	// 输出为文件
	out, err := os.Create("template.test.html")

	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}
