FROM golang:1.22.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /go/bin/app .

FROM golang:1.22.1-alpine

COPY --from=builder /go/bin/app /go/bin/app

EXPOSE 8080

CMD ["/go/bin/app"]