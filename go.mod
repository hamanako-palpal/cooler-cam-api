module github.com/hamanako-palpal/go_cooler_cam_api

go 1.14

require (
	github.com/gorilla/mux v1.7.4
	github.com/hamanako-palpal/go_cooler_cam_api/entities v0.0.0-00010101000000-000000000000 // indirect
	github.com/hamanako-palpal/go_cooler_cam_api/handle v0.0.0-00010101000000-000000000000
	github.com/hamanako-palpal/go_cooler_cam_api/infra v0.0.0-00010101000000-000000000000 // indirect
	github.com/hamanako-palpal/go_cooler_cam_api/repositories v0.0.0-00010101000000-000000000000 // indirect
	github.com/hamanako-palpal/go_cooler_cam_api/services v0.0.0-00010101000000-000000000000 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/lib/pq v1.7.0
	golang.org/x/net v0.0.0-20200625001655-4c5254603344 // indirect
	google.golang.org/api v0.28.0 // indirect
)

replace github.com/hamanako-palpal/go_cooler_cam_api/handle => ./handle

replace github.com/hamanako-palpal/go_cooler_cam_api/entities => ./entities

replace github.com/hamanako-palpal/go_cooler_cam_api/services => ./services

replace github.com/hamanako-palpal/go_cooler_cam_api/repositories => ./repositories

replace github.com/hamanako-palpal/go_cooler_cam_api/infra => ./infra
