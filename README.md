# cooler-cam-api
- Golang(1.14)
- air
- goose
- dotenv
# step
## build on env-develop
docker-compose build
## run docker
docker-compose up -d
## create table
docker-compose exec cooler-cam-api bash
goose -path go/src/work/db up
exit
## access
localhost:8000/api/smpl