package services

import (
	"encoding/json"
	"fmt"
	"go_cooler_cam_api/entities"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"

	// package名はvisionになっている
	"google.golang.org/api/vision/v1"
)

// CreateService サービスクライアント生成
func CreateService() *vision.Service {

	raw, err := ioutil.ReadFile("./VisionIdentify.json")
	if err != nil {
		fmt.Println("CreateService error0")
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

// Annotate 分析
func Annotate(cm *entities.Images) []byte {

	// 使いたいVisionの機能
	feature := &vision.Feature{
		Type:       "LABEL_DETECTION",
		MaxResults: 3,
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
	res, err := CreateService().Images.Annotate(batch).Do()
	if err != nil {
		fmt.Println("Annotate error1")
		log.Fatal(err)
		return nil
	}

	// 結果をJSON出力してみる
	body, _ := json.Marshal(res.Responses[0].LabelAnnotations)

	return body

}
