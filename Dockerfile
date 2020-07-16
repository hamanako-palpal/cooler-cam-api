# ベースとなるDockerイメージ指定
FROM golang:latest

ENV GO111MODULE=on
ENV EXEC_ENV=develop
# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/work
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/work
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/work

# air,goose,wireのインストール
RUN go get -u -v github.com/cosmtrek/air \
    github.com/google/wire/cmd/wire \
    bitbucket.org/liamstask/goose/cmd/goose

RUN go mod download

EXPOSE 8000