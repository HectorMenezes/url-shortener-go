FROM golang:1.20-alpine

RUN apk add postgresql

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./

RUN go install github.com/swaggo/swag/cmd/swag@latest && swag init
RUN CGO_ENABLED=0 GOOS=linux go build -o /url-shortener

CMD ["/url-shortener"]

ENTRYPOINT ["/bin/sh", "/app/entrypoint.sh"]
