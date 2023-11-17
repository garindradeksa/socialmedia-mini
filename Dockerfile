FROM golang:1.21-alpine

WORKDIR /app

COPY local.env .

RUN go build -o app .

CMD ["/app/app"]