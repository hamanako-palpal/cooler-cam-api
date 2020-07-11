package handle

import (
	"github.com/hamanako-palpal/cooler-cam-api/entities"
	"github.com/hamanako-palpal/cooler-cam-api/infra"
	"github.com/hamanako-palpal/cooler-cam-api/services"

	"encoding/json"
	"fmt"
	"net/http"
)

// CamHandler カメラ
type CamHandler struct {
	camservice *services.CamService
}

// InitCamHandler 初期化
func InitCamHandler() *CamHandler {

	visioncli := infra.InitVisionCli()
	labelcli := infra.InitLabelCli()
	camService := services.NewCamService(visioncli, labelcli)

	return &CamHandler{
		camservice: camService,
	}
}

// SmplHandler 導通確認
func SmplHandler(w http.ResponseWriter, r *http.Request) {

	img := &entities.ImageRequest{FileName: "aa", Contents: "smpl"}
	res, _ := json.Marshal(img)
	w.Write(res)
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
