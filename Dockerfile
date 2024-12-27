FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV DOCKERIZE_VERSION v0.9.1

RUN apt-get update \
    && apt-get install -y wget \
    && wget -O - https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz | tar xzf - -C /usr/local/bin \
    && apt-get autoremove -yqq --purge wget && rm -rf /var/lib/apt/lists/*

RUN GOS=linux CGO_ENABLED=0 go build -o ordersystem ./cmd/ordersystem/main.go ./cmd/ordersystem/wire_gen.go

COPY cmd/ordersystem/.env .env

EXPOSE 8080
EXPOSE 8000
EXPOSE 50051

CMD ["./ordersystem"]
