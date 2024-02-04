FROM golang:1.19.2-alpine

RUN apk update && apk --no-cache add git build-base

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN go install golang.org/x/tools/gopls@latest && \ 
  go install github.com/rubenv/sql-migrate/...@latest && \
  go install github.com/cosmtrek/air@latest && \
  go install github.com/golang/mock/mockgen@latest && \
  go install github.com/swaggo/swag/cmd/swag@latest

CMD ["air"]