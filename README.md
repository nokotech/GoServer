# Go Server

## run server
* `go run src/server.go`

## run docker
* `docker-compose up`

## build server
* `GOOS=windows GOARCH=386 go build -o build/server.exe src/server.go`
* `GOOS=linux GOARCH=amd64 go build -o build/server src/server.go`
