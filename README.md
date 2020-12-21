# cooler-cam-api
- golang
- air
- goose
- dotenv
- gorilla web-toolkit
- wire
# How to run?
## download tools
- docker
- docker-compose
- git
## Step
```
git clone https://github.com/hamanako-palpal/cooler-cam-api.git
cd cooler-cam-api
```
## get Identify-Key(Google-Vision-Api)
place file "VisionIdentify.json" on current directry(cooler-cam-api/VisionIdentify.json)
select pls.
- get file uploaded to Trello
- get your GCP account 'n' issue yourself
## build on env-develop
`docker-compose build`
## run docker
`docker-compose up -d`
## create table
```
docker-compose exec app bash
$ goose -path go/src/work/db up  
$ exit
```
## access
```
open localhost:8000/api/smpl
```
## finish?
```
docker-compose down
```
# product Env
```
docker build -f docker/Dockerfile.deploy -t heroku-deploy .
docker run --rm -it heroku-deploy
$ cd cooler-cam-api
$ heroku login -i
set mail-address
set password

$ git config --global user.email "<your-account>"
$ git config --global user.name "<youe-name>"
$ heroku git:remote -a new-wave-cooler-app
$ git add .
$ git commit -am "something message"
$ heroku stack:set container
$ git push heroku master
```