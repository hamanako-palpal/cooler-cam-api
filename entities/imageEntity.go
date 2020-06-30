package entities

// ImageRequest filename: ファイル名、contents: base64にencodeされた画像
type ImageRequest struct {
	FileName string `json:filename`
	Contents string `json:contents`
}
