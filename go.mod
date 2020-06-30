module github.com/hamanako-palpal/cooler-cam-api

go 1.14

require (
	github.com/gorilla/mux v1.7.4
	github.com/hamanako-palpal/cooler-cam-api/entities v0.0.0-00010101000000-000000000000 // indirect
	github.com/hamanako-palpal/cooler-cam-api/handle v0.0.0-00010101000000-000000000000
	github.com/hamanako-palpal/cooler-cam-api/infra v0.0.0-00010101000000-000000000000 // indirect
	github.com/hamanako-palpal/cooler-cam-api/repositories v0.0.0-00010101000000-000000000000 // indirect
	github.com/hamanako-palpal/cooler-cam-api/services v0.0.0-00010101000000-000000000000 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/lib/pq v1.7.0
	golang.org/x/net v0.0.0-20200625001655-4c5254603344 // indirect
	google.golang.org/api v0.28.0 // indirect
)

replace github.com/hamanako-palpal/cooler-cam-api/handle => ./handle

replace github.com/hamanako-palpal/cooler-cam-api/entities => ./entities

replace github.com/hamanako-palpal/cooler-cam-api/services => ./services

replace github.com/hamanako-palpal/cooler-cam-api/repositories => ./repositories

replace github.com/hamanako-palpal/cooler-cam-api/infra => ./infra
