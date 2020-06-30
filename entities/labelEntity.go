package entities

// LabelModel ラベルテーブルのマッパ
type LabelModel struct {
	Label string
	Date  string
}

// LabelResponce VisionAPIのレスポンス
type LabelResponce struct {
	Description string
	Mid         string
	Score       float32
	Topicality  float32
}
