# stage 0

FROM golang:1.18.3-alpine3.16

WORKDIR /app

COPY . .

RUN apk add --no-cache gcc musl-dev

RUN go mod download

RUN go mod verify

RUN go build -v -o server

# stage 1

FROM alpine:3.16.0

WORKDIR /

COPY --from=0 /app/server .

EXPOSE 8080

ENTRYPOINT ["/server"]
