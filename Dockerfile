FROM golang:1.17-alpine as builder

WORKDIR /app
RUN apk add build-base

COPY go.mod go.sum ./

RUN go mod download

COPY . ./
RUN go build -o ./webhook ./cmd/

FROM golang:1.17-alpine
WORKDIR /app
COPY --from=builder /app/webhook /app/webhook
EXPOSE 8080

ENTRYPOINT [ "/app/webhook" ]
