package repositories

import (
	"github.com/hamanako-palpal/cooler-cam-api/entities"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"

	// package名はvisionになっている
	"google.golang.org/api/vision/v1"
)

// VisionCl 呼び出し用構造体
type VisionCl struct {
	service *vision.Service
}

// Create クライアント生成
func Create() *vision.Service {

	raw, err := ioutil.ReadFile("../VisionIdentify.json")
	if err != nil {
		fmt.Println("Not found identify key")
		log.Fatal(err)
		return nil
	}

	cfg, err := google.JWTConfigFromJSON(raw, vision.CloudPlatformScope)

	if err != nil {
		fmt.Println("CreateService error1")
		log.Fatal(err)

		return nil
	}

	// OAuth2の認可を付与したHTTPクライアントを作る
	client := cfg.Client(context.Background())
	// Vision APIのサービスを作る
	svc, err := vision.New(client)

	if err != nil {
		fmt.Println("CreateService error2")
		log.Fatal(err)
		return nil
	}
	return svc
}

// CallAnnotate VISION_API_LABEL_DETECTIONをコール
func (vcl *VisionCl) CallAnnotate(cm *entities.ImageRequest) []entities.LabelResponce {

	var lr []entities.LabelResponce

	// 使いたいVisionの機能
	feature := &vision.Feature{
		Type:       "LABEL_DETECTION",
		MaxResults: 10,
	}

	// 目当ての画像データ
	img := &vision.Image{Content: cm.Contents}

	req := &vision.AnnotateImageRequest{
		Image:    img,
		Features: []*vision.Feature{feature},
	}

	// 1回の呼び出しで複数の処理を要求できる
	batch := &vision.BatchAnnotateImagesRequest{
		Requests: []*vision.AnnotateImageRequest{req},
	}

	// 実際のAPIコールを実行
	res, err := vcl.service.Images.Annotate(batch).Do()
	if err != nil {
		fmt.Println("API error")
		log.Fatal(err)
		return nil
	}
	// 結果をJSON出力してみる
	body, _ := json.Marshal(res.Responses[0].LabelAnnotations)
	err = json.Unmarshal(body, &lr)
	if err != nil {
		fmt.Println("Can not mapping")
		log.Fatal(err)
		return nil
	}

	return lr
}