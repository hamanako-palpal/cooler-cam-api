package services

import (
	"time"

	"github.com/hamanako-palpal/cooler-cam-api/entities"
	"github.com/hamanako-palpal/cooler-cam-api/repositories"
)

// CamService カメラから受け取った画像を分析
type CamService struct {
	vrepo repositories.VisionRepository
	lrepo repositories.LabelRepository
}

// NewCamService 初期化
func NewCamService(v repositories.VisionRepository, l repositories.LabelRepository) *CamService {

	return &CamService{
		vrepo: v,
		lrepo: l,
	}
}

// GetLabel 画像のラベルの取得
func (cs CamService) GetLabel(cm *entities.ImageRequest) []entities.LabelResponce {

	bd := cs.vrepo.CallAnnotate(cm)

	return bd
}

// SelectHighScore 好スコアの項目のみ抽出する(90%以上)
func (cs CamService) SelectHighScore(lres []entities.LabelResponce) []entities.LabelModel {

	var selected []entities.LabelModel
	for i := 0; i < len(lres); i++ {

		if lres[i].Score > 90.0 {
			selected = append(selected, entities.LabelModel{
				Label: lres[i].Description,
				Date:  time.Now().String(),
			})
		}
	}
	return selected
}

// InsertLabels ラベルテーブルへの挿入
func (cs CamService) InsertLabels(lms []entities.LabelModel) *entities.Request {

	for i := 0; i < len(lms); i++ {

		req := cs.lrepo.InsertOne(&lms[i])
		if req.Status != 500 {
			return &entities.Request{
				Status: 400,
				Result: "ng",
			}
		}
	}
	return &entities.Request{
		Status: 500,
		Result: "ok",
	}
}

// GetAllLabels ラベルテーブル全件取得
func (cs CamService) GetAllLabels() []entities.LabelModel {

	return cs.lrepo.FindAll()
}
