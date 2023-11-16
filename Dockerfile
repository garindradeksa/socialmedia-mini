FROM golang:1.21-alpine

COPY local.env /app

WORKDIR /app

RUN go mod tidy

RUN go build -o app .

CMD ["/app/app"]