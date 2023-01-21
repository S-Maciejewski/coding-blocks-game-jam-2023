FROM golang:1.19.5

RUN apt-get update && apt-get install -y \
    xorg-dev

WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
#ENV GOOS=windows GOARCH=amd64
ENV GOOS=js GOARCH=wasm
CMD go build -o /usr/local/bin/build/ebiten_2048.wasm
