FROM golang:1.17-alpine

RUN apk add git
COPY . /
RUN go mod download
WORKDIR /

ENTRYPOINT ["go", "run", "/github.go"]
