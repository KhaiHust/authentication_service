FROM golang:1.21-alpine AS builder
COPY ./src /go/src
WORKDIR /go/src/migration
RUN env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./migration \
    -ldflags="-X" \

FROM golang:1.21-alpine
COPY --from=builder /go/src/migration/config /app/migration
COPY --from=builder /go/src/migration/migration /app
EXPOSE 8080
WORKDIR /app
RUN chmod +x ./app
CMD [ "./migration"]