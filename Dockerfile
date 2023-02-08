# компилируем приложение
FROM golang:alpine as builder
WORKDIR /api-proxy
RUN apk --no-cache add ca-certificates
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./proxy-service

# запускаем собранное приложение
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /api-proxy/proxy-service /proxy-service
ENTRYPOINT ["/proxy-service"]