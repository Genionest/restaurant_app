
FROM golang:1.23.5-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOPROXY=https://goproxy.cn,direct \
    GOPATH=/go

WORKDIR /backend

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["go", "run", "."]