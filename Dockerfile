FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
COPY . .

RUN apk add make

RUN make build

EXPOSE 8000

CMD ["./bin/api"]