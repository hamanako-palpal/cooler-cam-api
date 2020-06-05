package entities

// Request api実行結果
type Request struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}
