FROM golang:1.24-bookworm AS base

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download
RUN go mod tidy

COPY . .

RUN go build -o go-project

EXPOSE 8000

CMD ["/build/go-project"]

