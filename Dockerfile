FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o pocketbase-app

ENTRYPOINT [ "/app/pocketbase-app", "serve", "--http", "0.0.0.0:8090" ]
