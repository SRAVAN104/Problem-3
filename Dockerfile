# Dockerfile
FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod init datetime-webapp 2>/dev/null || true

COPY . .

RUN go build -o /datetime-webapp

EXPOSE 8080

CMD ["/datetime-webapp"]