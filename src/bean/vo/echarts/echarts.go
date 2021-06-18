package echarts

// 键值对数据
type NVData struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

// 饼状图
type PieChart struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// 饼状图
type PieChartInt struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

// 折线图
type LineChart struct {
	XAxis  `json:"xAxis"`
	YAxis  `json:"yAxis"`
	Series []Series `json:"series"`
}

// 数据
type Series struct {
	Name string        `json:"name"`
	Data []interface{} `json:"data"`
}

// X 轴
type XAxis struct {
	Data []string `json:"data"`
}

// Y 轴
type YAxis struct {
	Data []string `json:"data"`
}

// 柱状图
type Histogram struct {
	XAxis  `json:"xAxis"`
	YAxis  `json:"yAxis"`
	Series []Series `json:"series"`
}
