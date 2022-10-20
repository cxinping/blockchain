package parameter

const (
	HLTV_INDEX       = "https://www.hltv.org"         // HLTV 首页地址
	MATCH_RESULT_URL = "https://www.hltv.org/results" //已经结束的比赛结果地址

	// 赛程/赛果
	TT_MATCH_URL           = "https://www.hltv.org/matches" // 赛事和赛程网页地址
	MATCH_STATUS_UNSTARTED = "unstarted"                    // 比赛未开始
	MATCH_STATUS_LIVE      = "live"                         // 比赛正在进行中
	MATCH_STATUS_OVER      = "over"                         // 比赛结束
	MATCH_MODE_ONLINE      = "online"                       // 线上
	MATCH_MODE_LAN         = "lan"                          // 线下

	// 错误信息
	ERROR_STATUS_UNDO = "undo" //未处理
	ERROR_STATUS_DONE = "dong" //已处理
)
