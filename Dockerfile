# ベースとなるDockerイメージ指定
FROM golang:latest

ENV GO111MODULE=on
ENV EXEC_ENV=product
# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/work
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/work
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/work

# airのインストール
RUN go get -u github.com/cosmtrek/air \
    bitbucket.org/liamstask/goose/cmd/goose

RUN go mod download

EXPOSE 8000