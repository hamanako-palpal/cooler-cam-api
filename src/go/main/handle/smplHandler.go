package handle

import (
	"encoding/json"
	"net/http"

	"github.com/hamanako-palpal/cooler-cam-api/src/go/main/entities"
)

// SmplHandler 導通チェック用
type SmplHandler struct{}

// InitSmplHandler 初期化
func InitSmplHandler() *SmplHandler {
	return &SmplHandler{}
}

// Ping 導通確認
func (s *SmplHandler) Ping(w http.ResponseWriter, r *http.Request) {
	img := &entities.ImageRequest{FileName: "aa", Contents: "smpl-test"}
	res, _ := json.Marshal(img)
	w.Write(res)
}
