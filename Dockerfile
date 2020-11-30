FROM golang:alpine

COPY . /project

WORKDIR /project

RUN go build -o build/go_main_service ./cmd/apiserver/main.go

CMD ["/project/build/go_main_service"]