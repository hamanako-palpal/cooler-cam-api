package handle

import (
	"github.com/hamanako-palpal/cooler-cam-api/src/go/main/entities"
	"github.com/hamanako-palpal/cooler-cam-api/src/go/main/services"

	"encoding/json"
	"fmt"
	"net/http"
)

// CamHandler カメラ
type CamHandler struct {
	camservice *services.CamService
}

// InitCamHandler 初期化
func InitCamHandler(camsvc *services.CamService) *CamHandler {
	return &CamHandler{
		camservice: camsvc,
	}
}

// AnnotateImage 受信したカメラの画像をvisionAPIに投入
func (camhdlr *CamHandler) AnnotateImage(w http.ResponseWriter, r *http.Request) {

	// リクエストマッピング
	var img entities.ImageRequest
	err := json.NewDecoder(r.Body).Decode(&img)

	if err != nil {
		fmt.Println("CamHandler error: MISS MAPPING")
		return
	}

	visionRes := camhdlr.camservice.GetLabel(&img)
	selectedRecord := camhdlr.camservice.SelectHighScore(visionRes)
	responce := camhdlr.camservice.InsertLabels(selectedRecord)
	resjson, _ := json.Marshal(responce)

	w.Header().Set("Content-Type", "application/json")
	w.Write(resjson)

}

// ViewAllLabels 全件取得
func (camhdlr *CamHandler) ViewAllLabels(w http.ResponseWriter, r *http.Request) {
	responce := camhdlr.camservice.GetAllLabels()
	resjson, _ := json.Marshal(responce)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resjson)
}
