FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/site main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/site /app/site
COPY --from=builder /app/html /app/html
COPY --from=builder /app/css /app/css
EXPOSE 8080
CMD ["/app/site"]
