FROM golang:1.26.4-alpine3.24 AS builder
WORKDIR /src
COPY . .
RUN GOOS=linux go build -a -o app .

FROM alpine:3.24
WORKDIR /root/
COPY --from=builder /src/app .
CMD ["./app"] 