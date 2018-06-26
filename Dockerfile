FROM golang:1.10.3-alpine3.7

MAINTAINER Vitali Makarov <405112@gmail.com>

ENV PROJECT_PATH=/go/src/github.com/clevertechru/simple-proxy

WORKDIR $PROJECT_PATH

COPY . .

RUN apk upgrade --update --no-cache && \
    apk add --update --no-cache bash curl git

RUN ./install.sh
RUN go build -ldflags "-X github.com/clevertechru/simple-proxy/handler.Buildstamp=`date -u +%Y/%m/%d_%H:%M:%S` -X github.com/clevertechru/public-api/handler.Commit=`git rev-parse HEAD`"

CMD ./simple-proxy

EXPOSE 3000