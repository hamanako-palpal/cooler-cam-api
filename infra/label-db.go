package infra

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/hamanako-palpal/cooler-cam-api/entities"
)

// LabelCli ラベルテーブルクライアント
type LabelCli struct {
	db *sql.DB
}

// InitLabelCli クライアント作成
func InitLabelCli() LabelCli {
	dburi := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_INST"),
		os.Getenv("DB_SSLMODE"))

	db, err := sql.Open("postgres", dburi)
	defer db.Close()

	if err != nil {
		fmt.Println(err)
		return LabelCli{
			db: nil,
		}
	}
	return LabelCli{
		db: db,
	}
}

// func (cli labelCli)Find(key string) entities.LabelModel {
//
// }

// FindAll 全件取得
func (cli LabelCli) FindAll() []entities.LabelModel {

	var lmodels []entities.LabelModel

	rows, err := cli.db.Query("SELECT * FROM labels;")
	if err != nil {
		fmt.Println(err)
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
func (cli LabelCli) InsertOne(lm *entities.LabelModel) *entities.Request {

	// INSERT
	insertQuery := "insert into labels (label, date) values ($1, $2);"
	_, err := cli.db.Exec(insertQuery, lm.Label, lm.Date)

	if err != nil {
		fmt.Println(err)
		return &entities.Request{
			Status: 400,
			Result: "ng",
		}
	}
	return &entities.Request{
		Status: 500,
		Result: "ok",
	}
}
