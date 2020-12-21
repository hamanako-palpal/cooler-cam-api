module github.com/hamanako-palpal/cooler-cam-api

go 1.15

require (
	github.com/google/wire v0.4.0
	github.com/gorilla/mux v1.8.0
	github.com/hamanako-palpal/cooler-cam-api/src/go/main/config v0.0.0-00010101000000-000000000000
	github.com/hamanako-palpal/cooler-cam-api/src/go/main/entities v0.0.0-00010101000000-000000000000 // indirect
	github.com/hamanako-palpal/cooler-cam-api/src/go/main/handle v0.0.0-00010101000000-000000000000
	github.com/hamanako-palpal/cooler-cam-api/src/go/main/infra v0.0.0-00010101000000-000000000000
	github.com/hamanako-palpal/cooler-cam-api/src/go/main/repositories v0.0.0-00010101000000-000000000000 // indirect
	github.com/hamanako-palpal/cooler-cam-api/src/go/main/services v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.3.0
	github.com/lib/pq v1.9.0
	golang.org/x/net v0.0.0-20200625001655-4c5254603344 // indirect
	google.golang.org/api v0.28.0 // indirect
)

replace github.com/hamanako-palpal/cooler-cam-api/src/go/main/handle => ./src/go/main/handle

replace github.com/hamanako-palpal/cooler-cam-api/src/go/main/entities => ./src/go/main/entities

replace github.com/hamanako-palpal/cooler-cam-api/src/go/main/services => ./src/go/main/services

replace github.com/hamanako-palpal/cooler-cam-api/src/go/main/repositories => ./src/go/main/repositories

replace github.com/hamanako-palpal/cooler-cam-api/src/go/main/infra => ./src/go/main/infra

replace github.com/hamanako-palpal/cooler-cam-api/src/go/main/config => ./src/go/main/config
