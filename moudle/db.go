package moudle

type RequestData struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type Fenlei_data struct {
	Yaowen    []int    `json:"yaowen"`
	Dangzheng []int    `json:"dangzheng"`
	Guandian  []int    `json:"guandian"`
	Difang    []int    `json:"difang"`
	Qita      []int    `json:"qita"`
	Date      []string `json:"date"`
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

type TimeVlues struct {
	StartDates  []string `json:"start_dates"`
	EndDates    []string `json:"end_dates"`
	StartValues []int    `json:"start_values"`
	EndValues   []int    `json:"end_values"`
}

type TypeCount struct {
	Timestamp string `xorm:"timestamp"`
	Count     int    `xorm:"count"`
}
