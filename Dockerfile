FROM golang:1.15-alpine

RUN apk add git
RUN  go get -u github.com/go-telegram-bot-api/telegram-bot-api \
    && go get github.com/google/go-github/github \
    && go get github.com/olekukonko/tablewriter

COPY github.go /

ENTRYPOINT ["go", "run", "/github.go"]
