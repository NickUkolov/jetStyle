FROM golang:1.22-alpine as builder
RUN apk add upx --no-cache && \
    mkdir /source
WORKDIR /source
COPY . .
RUN go mod download && go build -o ./notes-service cmd/main.go && upx -9 ./notes-service

FROM alpine:3
RUN addgroup -g 1000 notroot && \
    adduser -D -u 1000 -G notroot notroot

RUN mkdir /app

COPY --from=builder source/notes-service /app/notes-service

RUN chmod +x /app/notes-service && \
    chown -R notroot:notroot /app

USER notroot
WORKDIR /app
EXPOSE $NOTES_PORT
CMD ["./notes-service"]
