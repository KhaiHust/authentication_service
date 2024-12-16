FROM golang:1.21-alpine AS builder
ARG BUILD_VERSION
ARG BUILD_COMMIT_HASH
ARG BUILD_TIME
ARG BS_PKG=github.com/KhaiHust/authen_service/public/bootstrap
ENV GO111MODULE=on
COPY ./src /go/src
WORKDIR /go/src/public
RUN env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./public \
    -ldflags="-X '$BS_PKG.Version=$BUILD_VERSION' -X '$BS_PKG.CommitHash=$BUILD_COMMIT_HASH' -X '$BS_PKG.BuildTime=$BUILD_TIME'"

FROM golang:1.21-alpine
COPY --from=builder /go/src/public/config /app/config
COPY --from=builder /go/src/public/public /app/public
EXPOSE 8080
WORKDIR /app
RUN chmod +x ./public
CMD ["./public"]