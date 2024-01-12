FROM golang:alpine
# RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base

WORKDIR /usr/app
COPY . .

RUN go install -mod=mod github.com/githubnemo/CompileDaemon
RUN go mod tidy

ENTRYPOINT CompileDaemon -build="go build -o m ./src/cmd/server" -command="./m"