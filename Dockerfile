FROM golang:1.24-alpine3.21 as builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
ENV CGO_ENABLED=1
RUN go build -o server .

FROM alpine

WORKDIR /
COPY --from=builder /database/migration ./database/migration

WORKDIR /app
COPY --from=builder /server .

EXPOSE 8080

CMD ["./server"]
