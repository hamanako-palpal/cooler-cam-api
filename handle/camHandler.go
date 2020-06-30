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
	cs *services.CamService
}

// InitCamHandler 初期化
func InitCamHandler() *CamHandler {

	vcl := infra.InitVisionCli()
	lcl := infra.InitLabelCli()
	camService := services.NewCamService(vcl, lcl)

	return &CamHandler{
		cs: camService,
	}
}

// SmplHandler 導通確認
func SmplHandler(w http.ResponseWriter, r *http.Request) {

	img := &entities.ImageRequest{FileName: "aa", Contents: "smpl"}
	res, _ := json.Marshal(img)
	w.Write(res)
}

// AnnotateImage 受信したカメラの画像をvisionAPIに投入
func (ch *CamHandler) AnnotateImage(w http.ResponseWriter, r *http.Request) {

	// リクエストマッピング
	var img entities.ImageRequest
	err := json.NewDecoder(r.Body).Decode(&img)

	if err != nil {
		fmt.Println("CamHandler error: MISS MAPPING")
		return
	}

	visionRes := ch.cs.GetLabel(&img)
	selectedRecord := ch.cs.SelectHighScore(visionRes)
	responce := ch.cs.InsertLabels(selectedRecord)
	resjson, _ := json.Marshal(responce)

	w.Header().Set("Content-Type", "application/json")
	w.Write(resjson)

}

// ViewAllLabels 全件取得
func (ch *CamHandler) ViewAllLabels(w http.ResponseWriter, r *http.Request) {
	responce := ch.cs.GetAllLabels()
	resjson, _ := json.Marshal(&responce)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resjson)
}
