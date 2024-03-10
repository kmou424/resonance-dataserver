# Builder Stage
FROM golang:latest as builder

WORKDIR /app

USER root

RUN apt-get update && apt-get install -y git

ARG ED25519_KEY

RUN mkdir -p /root/.ssh/ && \
    echo "$ED25519_KEY" > /root/.ssh/id_ed25519 && \
    chmod 600 /root/.ssh/id_ed25519 && \
    ssh-keyscan -t ed25519 github.com >> ~/.ssh/known_hosts

RUN git clone git@github.com:kmou424/resonance-dataserver.git /app/dataserver
RUN git clone git@github.com:kmou424/resonance-microservice.git /app/microservice

RUN cd /app/dataserver && \
    curl -L "https://raw.githubusercontent.com/kmou424/resonance-resodata/main/goods_cities_mapper.json" \
         > database/data/files/goods_cities_mapper.json && \
    go mod download && \
    CGO_ENABLED=1 go build -v -a -trimpath -ldflags "-s -w -extldflags '-static'" -o server_bin ./cmd/server

RUN cd /app/microservice && \
    go mod download && \
    CGO_ENABLED=0 go build -v -a -trimpath -ldflags "-s -w --extldflags '-static'" -o daemon ./cmd


FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache supervisor

COPY --from=builder /app/dataserver/server_bin /app/server
COPY --from=builder /app/microservice/daemon /app/daemon

COPY supervisord/* /etc/supervisord/
RUN chmod 644 /etc/supervisord/*

CMD ["supervisord", "-c", "/etc/supervisord/supervisord.conf", "-n"]
