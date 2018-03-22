FROM golang:1.7

ENV GOROOT /usr/local/go
RUN go get github.com/garyburd/redigo/redis
RUN cp -fr /go /usr/local


EXPOSE 3000
