package infra

import (
	"database/sql"
	"fmt"

	"github.com/hamanako-palpal/cooler-cam-api/src/go/main/entities"
	"github.com/hamanako-palpal/cooler-cam-api/src/go/main/repositories"
)

// LabelCli ラベルテーブルクライアント
type LabelCli struct {
	db *sql.DB
}

// InitLabelCli クライアント作成
func InitLabelCli(db *sql.DB) repositories.LabelRepository {

	return &LabelCli{
		db: db,
	}
}

// func (cli labelCli)Find(key string) entities.LabelModel {
//
// }

// FindAll 全件取得
func (cli *LabelCli) FindAll() []entities.LabelModel {

	var lmodels []entities.LabelModel

	rows, err := cli.db.Query("SELECT * FROM labels;")
	if err != nil {
		fmt.Println(err)
		return lmodels
	}
	if rows == nil {
		fmt.Println("row is nil!!!!!!!!!!!!!!!!")
		return lmodels
	}

	for rows.Next() {
		var l string
		var d string

		if err := rows.Scan(&l, &d); err != nil {
			fmt.Println(err)
		}
		lmodels = append(lmodels, entities.LabelModel{Label: l, Date: d})
	}
	return lmodels
}

// InsertOne 更新操作
func (cli *LabelCli) InsertOne(lm entities.LabelModel) *entities.Request {

	// INSERT
	insertQuery := "insert into labels (LABEL, DATE) values ($1, $2);"

	_, err := cli.db.Exec(insertQuery, lm.Label, lm.Date)

	if err != nil {
		fmt.Println(err)
		return &entities.Request{
			Status: 400,
			Result: "miising InsertOne",
		}
	}
	return &entities.Request{
		Status: 200,
		Result: "ok",
	}
}
