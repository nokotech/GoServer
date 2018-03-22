# Go Server

## get lib
* `go get github.com/garyburd/redigo/redis`

## run server
* `go run src/server.go`

## run docker
* `docker build .docker/app`
* `docker-compose up`
* `docker-compose down`

## build server
* `GOOS=windows GOARCH=386 go build -o build/server.exe src/server.go`
* `GOOS=linux GOARCH=amd64 go build -o build/server src/server.go`
