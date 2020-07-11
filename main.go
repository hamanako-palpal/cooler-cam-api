package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hamanako-palpal/cooler-cam-api/handle"
	_ "github.com/lib/pq"

	"fmt"

	"github.com/joho/godotenv"
)

func main() {

	// EXEC_ENVはDockerfile内で指定
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("EXEC_ENV")))
	if err != nil {
		fmt.Println("CANNNOT READ ENV FILE")
	}

	fmt.Println("Server Open!")
	camhandler := handle.InitCamHandler()

	r := mux.NewRouter()
	r.HandleFunc("/api/smpl", handle.SmplHandler).Methods("GET")
	r.HandleFunc("/api/cam", camhandler.AnnotateImage).Methods("POST")
	r.HandleFunc("/api/cam/labels", camhandler.ViewAllLabels).Methods("GET")

	s := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: r,
	}
	s.ListenAndServe()
}
