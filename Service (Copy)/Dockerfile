FROM golang:1.22.2

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy
COPY . .

RUN go build -o main ./cmd

EXPOSE 8080

CMD ["./main"]
