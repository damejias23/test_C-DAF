FROM golang:1.10 AS build
WORKDIR /go/src
COPY go ./go
COPY main.go .
COPY go.sum .
COPY go.mod .

ENV CGO_ENABLED=0

RUN go get github.com/damejias23/test_C-DAF
RUN go build -o openapi .

FROM alpine:3.18 AS runtime
COPY --from=build /go/src/openapi ./
COPY .env ./
EXPOSE 8080/tcp 8882/tcp 3000

ENTRYPOINT ["./openapi"]
