FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod .  
RUN go mod download

COPY . ./  

RUN go build -o main cmd/server/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]
