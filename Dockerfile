FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/binary

FROM scratch

COPY --from=builder /app/binary /app/binary

ENTRYPOINT ["/app/binary"]