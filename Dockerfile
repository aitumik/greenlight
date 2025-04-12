FROM golang:1.18.2-alpine

WORKDIR /build

COPY go.sum .
COPY go.mod .

COPY . .

RUN go mod download

RUN go build -o greenlight cmd/api/*

CMD ["./greenlight"]

