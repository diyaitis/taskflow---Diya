FROM golang:1.21 as builder
WORKDIR /app
COPY . .
RUN go build -o main ./cmd

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
