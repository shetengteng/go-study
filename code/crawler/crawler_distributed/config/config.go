package config

const (
	// service ports
	ItemSaverPort = 1234
	WorkerPort0   = ":9000"
	WorkerPort1   = ":9001"

	// ElasticSearch
	ElasticIndex = "dating_profile"

	// RPC Endpoints
	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"

	// Parser Names
	ParseCityList = "ParseCityList"
	ParseCity     = "ParseCity"
	NilParser     = "NilParser"
	ProfileParser = "ProfileParser"
)
