package view

import (
	"html/template"
	"io"
	"learngo/crawler/frontend/model"
)

type SearchResultView struct {
	template *template.Template
}

func CreateSearchResultView(filename string) SearchResultView {
	return SearchResultView{
		// 返回带有模板的视图
		template.Must(template.ParseFiles(filename)),
	}
}

func (s SearchResultView) Render(w io.Writer, data model.SearchResult) error {
	// 将数据给视图模板，并给模板中的变量赋值，最后输出到w中
	return s.template.Execute(w, data)
}
