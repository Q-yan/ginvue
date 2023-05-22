package moudle

type RequestData struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type Fenlei_data struct {
	Yaowen    []int `json:"yaowen"`
	Dangzheng []int `json:"dangzheng"`
	Guandian  []int `json:"guandian"`
	Difang    []int `json:"difang"`
	Qita      []int `json:"qita"`
}

type Item struct {
	Date []string `json:"date"`
	Num  []int    `json:"num"`
	Bar  []int    `json:"bar"`
	Line []int    `json:"line"`
}

type Ites struct {
	Date string `json:"date"`
	Num  int    `json:"num"`
	Bar  int    `json:"bar"`
	Line int    `json:"line"`
}

type ll struct {
}

type Liuliang struct {
	Timestamp             string `json:"timestamp"`
	ResponseContentLength int    `json:"response_content_length"`
}
