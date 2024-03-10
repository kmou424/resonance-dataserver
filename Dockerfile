FROM golang:latest AS builder

WORKDIR /app

ENV GOPROXY https://goproxy.cn,direct

COPY . /app

RUN curl -L "https://mirror.ghproxy.com/https://raw.githubusercontent.com/kmou424/resonance-resodata/main/goods_cities_mapper.json" \
     > database/data/files/goods_cities_mapper.json || exit 1

RUN go mod download
RUN CGO_ENABLED=1 go build -v=0 -a -trimpath -ldflags "-s -w -extldflags '-static'" -o server_bin ./cmd/server

FROM scratch
WORKDIR /app
COPY --from=builder /app/server_bin /app/server
CMD ["/app/server"]
