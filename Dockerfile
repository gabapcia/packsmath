FROM golang:1.24 AS builder

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o packsmath cmd/api/main.go

FROM scratch

COPY --from=builder /app/packsmath /packsmath

EXPOSE 3000

ENTRYPOINT ["/packsmath"]
