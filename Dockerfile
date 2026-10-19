FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/app /app/app
COPY --from=builder /app/fundo.png.jpg /app/fundo.png.jpg

ENV HOST=0.0.0.0
ENV PORT=8080

EXPOSE 8080

CMD ["./app"]
