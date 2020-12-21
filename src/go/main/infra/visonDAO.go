package infra

import (
	"encoding/json"

	"github.com/hamanako-palpal/cooler-cam-api/src/go/main/entities"
	"github.com/hamanako-palpal/cooler-cam-api/src/go/main/repositories"

	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"

	// package名はvisionになっている
	"google.golang.org/api/vision/v1"
)

// VisionCli 呼び出し用構造体
type VisionCli struct {
	svc *vision.Service
}

// InitVisionCli クライアント生成
func InitVisionCli() repositories.VisionRepository {

	raw, err := ioutil.ReadFile("../../../../../VisionIdentify.json")
	if err != nil {
		fmt.Println("Not found identify key")
		log.Fatal(err)
		return &VisionCli{
			svc: nil,
		}
	}

	cfg, err := google.JWTConfigFromJSON(raw, vision.CloudPlatformScope)

	if err != nil {
		fmt.Println("Initialize error: CANNOT READ IDENTIFY-KEY-FILE")
		log.Fatal(err)

		return &VisionCli{
			svc: nil,
		}
	}

	// OAuth2の認可を付与したHTTPクライアントを作る
	client := cfg.Client(context.Background())
	// Vision APIのサービスを作る
	svc, err := vision.New(client)

	if err != nil {
		fmt.Println("Initialize error: CANNOT CREATE API-SERVICE")
		log.Fatal(err)
		return &VisionCli{
			svc: nil,
		}
	}
	return &VisionCli{
		svc: svc,
	}
}

// CallAnnotate VISION_API_LABEL_DETECTIONをコール
func (vcli VisionCli) CallAnnotate(ir *entities.ImageRequest) []entities.LabelResponce {

	// 初期化チェック
	if vcli.svc == nil {
		fmt.Println("CallAnnotate error: NOT INITIALIZE")
	}

	// 使いたいVisionの機能
	feature := &vision.Feature{
		Type:       "LABEL_DETECTION",
		MaxResults: 10,
	}

	// 目当ての画像データ
	img := &vision.Image{Content: ir.Contents}

	req := &vision.AnnotateImageRequest{
		Image:    img,
		Features: []*vision.Feature{feature},
	}

	// 1回の呼び出しで複数の処理を要求できる
	batch := &vision.BatchAnnotateImagesRequest{
		Requests: []*vision.AnnotateImageRequest{req},
	}

	// 実際のAPIコールを実行
	res, err := vcli.svc.Images.Annotate(batch).Do()
	if err != nil {
		fmt.Println("API error")
		log.Fatal(err)
		return nil
	}

	var lr []entities.LabelResponce

	// 結果をJSON出力してみる
	body, _ := json.Marshal(res.Responses[0].LabelAnnotations)

	err = json.Unmarshal(body, &lr)
	if err != nil {
		fmt.Println("Can not mapping")
		log.Fatal(err)
		return nil
	}
	fmt.Println(len(lr))

	return lr
}
