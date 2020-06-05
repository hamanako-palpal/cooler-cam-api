package main

import (
	"fmt"
	"go_cooler_cam_api/handle"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Server Open!")
	r := mux.NewRouter()
	r.HandleFunc("/api/smpl", handle.SmplHandler).Methods("GET")
	r.HandleFunc("/api/cam", handle.CamHandler).Methods("POST")
	s := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	s.ListenAndServe()
	fmt.Println("Server state!")
}
