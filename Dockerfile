# syntax=docker/dockerfile:1

FROM golang:1.21.5
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /app/go-bookstore ./cmd/api
EXPOSE 8080
CMD ["/app/go-bookstore"]