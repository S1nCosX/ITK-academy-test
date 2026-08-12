FROM golang:1.26.4-alpine3.24 AS base
WORKDIR /src
COPY . .
RUN go mod download

FROM base AS test
RUN go test -v ./...

FROM base AS build
RUN GOOS=linux go build -a -o app .

FROM alpine:3.24
WORKDIR /root/
COPY --from=build /src/app .
CMD ["./app"] 