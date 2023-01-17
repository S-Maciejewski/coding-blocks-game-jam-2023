FROM golang:1.19.5

RUN apt-get update && apt-get install -y \
    xorg-dev

WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
ENV GOOS=windows GOARCH=amd64
CMD go build -o /usr/local/bin/build/
