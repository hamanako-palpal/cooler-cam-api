package main

import (
	"net/http"
	"os"

	"fmt"

	"github.com/gorilla/mux"
	"github.com/hamanako-palpal/cooler-cam-api/config"
	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

func main() {

	// EXEC_ENVはDockerfile内で指定
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("EXEC_ENV")))
	if err != nil {
		fmt.Println("CANNNOT READ ENV FILE")
	}

	// DBインスタンス生成
	db := config.CreateDB()
	handlerGen, initErr := Initialize(db)
	if initErr != nil {
		fmt.Println("CANNOT DI")
	}

	fmt.Println("Server Open!")
	r := mux.NewRouter()
	r.HandleFunc("/api/smpl", handlerGen.Smplhandler.Ping).Methods("GET")
	r.HandleFunc("/api/cam", handlerGen.Camhandler.AnnotateImage).Methods("POST")
	r.HandleFunc("/api/cam/labels", handlerGen.Camhandler.ViewAllLabels).Methods("GET")

	s := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: r,
	}
	s.ListenAndServe()
	defer func() {
		db.Close()
		fmt.Println("Server Close!")
	}()
}
