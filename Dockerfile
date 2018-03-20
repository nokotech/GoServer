FROM alpine:3.4

WORKDIR /app
COPY sever /app
EXPOSE 3000

CMD ["./server"]
