FROM golang:latest AS builder

WORKDIR /app

ENV GOPROXY https://goproxy.cn,direct

COPY . /app

RUN go mod download
RUN CGO_ENABLED=0 go build -v=0 -a -trimpath -ldflags "-s -w -extldflags '-static'" -o server_bin ./cmd/server

FROM scratch
WORKDIR /app
COPY --from=builder /app/server_bin /app/server
CMD ["/app/server"]
