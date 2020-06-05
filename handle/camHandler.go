package handle

import (
	"encoding/json"
	"fmt"
	"go_cooler_cam_api/entities"
	"go_cooler_cam_api/services"
	"log"
	"net/http"
)

// SmplHandler 導通確認
func SmplHandler(w http.ResponseWriter, r *http.Request) {

	aso := &entities.Images{FileName: "aa", Contents: "t11t41t54t"}
	res, _ := json.Marshal(aso)
	w.Write(res)
}

// CamHandler 受信したカメラの画像をvisionAPIに投入
func CamHandler(w http.ResponseWriter, r *http.Request) {

	var img entities.Images
	err := json.NewDecoder(r.Body).Decode(&img)

	if err != nil {
		fmt.Println("error1")
		log.Fatal(err)
		return
	}

	res, err := json.Marshal(img)
	if err != nil {
		fmt.Println("error2")
		log.Fatal(err)
		return
	}
	fmt.Println(string(res))

	w.Header().Set("Content-Type", "application/json")
	visionRes := services.Annotate(&img)
	w.Write(visionRes)

}
