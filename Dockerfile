FROM golang:1.22.4-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
RUN go mod verify

COPY . .
RUN go build -v -o main ./cmd/app/main.go


CMD ["/app/main"]
