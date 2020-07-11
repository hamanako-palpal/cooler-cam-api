package entities

// LabelModel ラベルテーブルのマッパ
type LabelModel struct {
	Label string
	Date  string
}

// LabelResponce VisionAPIのレスポンス
type LabelResponce struct {
	Description string  `json:description`
	Mid         string  `json:mid`
	Score       float64 `json:score`
	Topicality  float64 `json:topicality`
}
