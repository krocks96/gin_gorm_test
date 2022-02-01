FROM golang:1.16.7

COPY src /go/src

WORKDIR /go/src

RUN apt-get update && apt-get install -y \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/* \
  && go get github.com/gin-gonic/gin \
  && go get github.com/jinzhu/gorm \
  && go get github.com/go-sql-driver/mysql

EXPOSE 3000