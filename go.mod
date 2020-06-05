module go_cooler_cam_api

go 1.14

require (
	github.com/gorilla/mux v1.7.4
	go_cooler_cam_api/entities v0.0.0-00010101000000-000000000000 // indirect
	go_cooler_cam_api/handle v0.0.0-00010101000000-000000000000
	go_cooler_cam_api/services v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/net v0.0.0-20200506145744-7e3656a0809f // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	google.golang.org/api v0.23.0 // indirect
)

replace go_cooler_cam_api/handle => ./handle

replace go_cooler_cam_api/entities => ./entities

replace go_cooler_cam_api/services => ./services
