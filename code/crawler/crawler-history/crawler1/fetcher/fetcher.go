package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"bufio"

	"log"

	"math/rand"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var browserHeader = [...]string{
	"Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.153 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:30.0) Gecko/20100101 Firefox/30.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.75.14 (KHTML, like Gecko) Version/7.0.3 Safari/537.75.14",
	"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0)",
	"Mozilla/5.0 (Windows; U; Windows NT 5.1; it; rv:1.8.1.11) Gecko/20071127 Firefox/2.0.0.11",
	"Opera/9.25 (Windows NT 5.1; U; en)",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
	"Mozilla/5.0 (compatible; Konqueror/3.5; Linux) KHTML/3.5.5 (like Gecko) (Kubuntu)",
	"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.8.0.12) Gecko/20070731 Ubuntu/dapper-security Firefox/1.5.0.12",
	"Lynx/2.8.5rel.1 libwww-FM/2.14 SSL-MM/1.4.1 GNUTLS/1.2.9",
	"Mozilla/5.0 (X11; Linux i686) AppleWebKit/535.7 (KHTML, like Gecko) Ubuntu/11.04 Chromium/16.0.912.77 Chrome/16.0.912.77 Safari/535.7",
	"Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:10.0) Gecko/20100101 Firefox/10.0 ",
}

func httpGet(url string) (resp *http.Response, err error) {
	client := &http.Client{}

	// 模拟浏览器
	request, err := http.NewRequest("GET", url, nil)

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(browserHeader))

	request.Header.Add("User-Agent", browserHeader[i])

	return client.Do(request)
}

// 给一个URL返回一个text
func Fetch(url string) ([]byte, error) {

	// 1s一次，防止过快请求
	time.Sleep(time.Second*1)

	resp, err := httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,
			//errors.New("") // 方式1
			fmt.Errorf("wrong status code: %d ", resp.StatusCode) // 方式2
	}

	reader := bufio.NewReader(resp.Body)
	e := determineEncoding(reader)

	utf8Reader := transform.NewReader(reader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

// 确定网页的编码
func determineEncoding(reader *bufio.Reader) encoding.Encoding {
	// 读取开头1024个字节，用于DetermineEncoding方法的鉴别编码格式
	//peek是对bufio.Reader来说的，而非原始的reader，对bufio.Reader而言是从头读
	bytes, err := reader.Peek(1024)
	if err != nil {
		//panic(err)
		log.Printf("Fetcher error:%v ", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
