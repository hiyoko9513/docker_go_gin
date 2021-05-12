FROM golang:1.15-alpine3.12

WORKDIR /go/app

COPY . ../../hiyoko

RUN set -eux \
  && apk update \
  && apk upgrade \
  && apk add --no-cache git curl 
  # && go get github.com/go-sql-driver/mysql 
  # && go get github.com/lib/pq 

# Ginインストール
RUN go get github.com/gin-gonic/gin
RUN go get github.com/codegangsta/gin
# Ginの起動
CMD ["gin","--port","8080"]
