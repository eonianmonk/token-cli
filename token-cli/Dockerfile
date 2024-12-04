FROM golang:alpine as builder

ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOOS=linux

WORKDIR /build
ADD go.mod ./
ADD go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /app/token-cli ./cmd/*.go

# will be passed as volume 
# COPY ../config.yaml ./

WORKDIR /app

ENTRYPOINT ["./token-cli"]