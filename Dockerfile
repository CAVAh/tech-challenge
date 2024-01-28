FROM golang:1.21.4

WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go build -o bin .

EXPOSE 8080

ENTRYPOINT ["/app/bin"]