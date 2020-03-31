package parser

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"regexp"
)

func TestParseCityList(t *testing.T) {

	//contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	// 从本地读取
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)

	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCitys := []string{
		"City:阿坝",
		"City:阿克苏",
		"City:阿拉善盟",
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d request but had %d ",
			resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s ", i, url, result.Requests[i].Url)
		}
	}

	for i, city := range expectedCitys {
		if result.Items[i].(string) != city {
			t.Errorf("expected city #%d: %s; but was %s ", i, city, result.Items[i].(string))
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d items but had %d ",
			resultSize, len(result.Items))
	}

}

func TestParseCity(t *testing.T) {

	client := &http.Client{}
	url := "http://album.zhenai.com/u/1308430399"

	// 模拟浏览器
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1")

	resp, err := client.Do(request)

	//resp, err := http.Get("https://album.zhenai.com/u/1308430399")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("error: status code", resp.StatusCode)
		return
	}

	// 不做处理
	all, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", all)
}

func TestParseProfile(t *testing.T) {
	var s string = `<div class="m-btn purple" data-v-8b1eac0c>离异</div>`
	var marriageRe = regexp.MustCompile(`<div [^>]*>([离异未婚丧偶]+)</div>`)

	match := marriageRe.FindAllStringSubmatch(s,-1)
	fmt.Println(match)

}