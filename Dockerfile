FROM golang:1.22.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /go/bin/app .

FROM scratch

COPY --from=builder /go/bin/app /go/bin/app
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["/go/bin/app"]