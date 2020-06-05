package entities

// Images filename: ファイル名、contents: base64にencodeされた画像
type Images struct {
	FileName string `json:filename`
	Contents string `json:contents`
}
