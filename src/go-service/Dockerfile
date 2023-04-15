FROM golang:1.20.1

WORKDIR /app

COPY . ./
RUN go mod download

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go-svc

EXPOSE 3099

CMD ["/go-svc"]