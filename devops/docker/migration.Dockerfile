FROM golang:1.21-alpine AS builder
COPY ./src /go/src
WORKDIR /go/src/migration
RUN env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./migration

FROM golang:1.21-alpine
COPY --from=builder /go/src/migration/config /app/config
COPY --from=builder /go/src/migration/migration /app/migration
EXPOSE 8080
WORKDIR /app
RUN chmod +x ./migration
CMD ["./migration"]