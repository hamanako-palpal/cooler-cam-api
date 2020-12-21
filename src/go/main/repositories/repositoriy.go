package repositories

import "github.com/hamanako-palpal/cooler-cam-api/src/go/main/entities"

// LabelRepository ラベル管理
type LabelRepository interface {
	// Find(string) entities.LabelModel
	FindAll() []entities.LabelModel
	InsertOne(entities.LabelModel) *entities.Request
}

// VisionRepository VisionAPIコール
type VisionRepository interface {
	CallAnnotate(*entities.ImageRequest) []entities.LabelResponce
}
