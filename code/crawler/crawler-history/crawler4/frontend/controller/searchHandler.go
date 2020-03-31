package controller

import (
	"learngo/crawler-history/crawler4history/crawler4/frontend/view"
	"net/http"

	"strconv"
	"strings"

	"context"
	"learngo/crawler-history/crawler4history/crawler4/engine"
	"learngo/crawler-history/crawler4history/crawler4/frontend/model"
	"reflect"

	"regexp"

	"gopkg.in/olivere/elastic.v5"
)

type SearchResultHandler struct {
	view   view.SearchResultView // 渲染的view
	client *elastic.Client       // 获取数据的es客户端
}

// 生成handler
func CreateSearchResultHandler(template string) SearchResultHandler {

	// handler包含es客户端
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	return SearchResultHandler{
		view:   view.CreateSearchResultView(template), // 配置view
		client: client,                                // 配置客户端
	}
}

// 实现net/http/server.go接口
func (handler SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	// 实现 localhost:8888/search?q=男 已购房&from=20

	// 获取查询参数
	q := strings.TrimSpace(req.FormValue("q"))

	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	//fmt.Fprintf(w, "q=%s, from=%d", q, from)

	// 查询结果
	pageInfo, err := handler.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	// 渲染结果
	handler.Render(w, pageInfo)
}

// 将查询的结果给视图渲染
func (handler SearchResultHandler) Render(w http.ResponseWriter, re model.SearchResult) {
	err := handler.view.Render(w, re)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// 查询结果
func (handler SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {

	var result model.SearchResult

	// 调用es进行查询
	resp, err := handler.client.
		Search("dating_profile").
		Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).
		From(from).Do(context.Background())

	if err != nil {
		return result, err
	}

	result.Query = q // 用于回显查询
	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	//fmt.Printf("%+v",resp)
	//fmt.Println(result.Items)
	// 如果Items的类型是engine.Items 则使用
	//for _,v := range resp.Each(reflect.TypeOf(engine.Item{})){
	//	item := v.(engine.Item)
	//	result.Items = append(result.Items,item)
	//}

	return result, nil
}

// 针对Payload.Age:(>30) 的类型的Es查询，在搜索框中输入Age:(>30) 达到相同的效果
func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}
