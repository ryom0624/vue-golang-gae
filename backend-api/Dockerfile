FROM golang:1.12.9-alpine3.10 as build

WORKDIR /go/app

COPY . .

RUN set -ex \
    && apk update \
# gcc libc-devはalpineでのgo-sqlite3のコンパイルに必要
    && apk add --no-cache sqlite git sqlite gcc libc-dev

# article.sqlite3がなければ生成して読み込み
RUN touch ./db/articles.sqlite3 \
    && sqlite3 ./db/articles.sqlite3 < ./db/init.sql

# バイナリファイルを生成しalpineにマルチステージビルド
RUN go build -o backend-api \
    && go get gopkg.in/urfave/cli.v2@master \
    && go get github.com/oxequa/realize

FROM alpine:3.10 as binary

WORKDIR /app

# buildイメージからコピー
COPY --from=build /go/app/backend-api .

RUN set -x \
    && addgroup go \
    && adduser -D -G go go \
    && chown -R go:go /app/backend-api

CMD ["./backend-api"]