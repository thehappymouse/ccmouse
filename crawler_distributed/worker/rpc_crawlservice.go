package worker

import "dali.cc/ccmouse/crawler/engine"

type CrawlService struct{}

// request 包含一个interface类弄的parse，无法在网络传输，所以要改造
func (CrawlService) Process(req Request, result *ParseResult) error {
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}
	engineResult, err := engine.Worker(engineReq)
	// 返回内容
	*result = SerializeResult(engineResult)

	return nil
}
