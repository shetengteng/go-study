package worker

import "learngo/crawler/engine"

type CrawlService struct{}

// 输入req，处理结果result返回
func (CrawlService) Process(
	//req engine.Request, // Request 中的Parser不能在网络中传递
	req Request,
	result *ParseResult,
) error {

	// 将请求解析为engineRequest
	engineRequest, err := DeserializeRequest(req)
	if err != nil {
		return err
	}
	// 对返回的结果解析成传输的result
	engineResult, err := engine.Worker(engineRequest)
	if err != nil {
		return err
	}
	*result = SerializeResult(engineResult)
	return nil
}
