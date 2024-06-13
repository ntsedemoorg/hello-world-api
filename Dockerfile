FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go get github.com/bradfitz/gomemcache/memcache && go build -o main .

EXPOSE 8080

CMD ["./main"]
