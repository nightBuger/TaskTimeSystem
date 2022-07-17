package webengine

type SqlResult struct {
	Msg         string      `json:"msg"`
	ResultSlice interface{} `json:"data"`
	RowCount    int         `json:"-"`
}

type SqlResultSlice struct {
	Msg         string      `json:"msg"`
	ResultSlice interface{} `json:"data"`
	RowCount    int         `json:"-"`
	FoundRows   int         `json:"total"`
}
