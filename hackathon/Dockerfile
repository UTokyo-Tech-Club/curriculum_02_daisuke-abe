FROM golang:1.19.2

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

CMD ["go", "run", "main.go" "transactions.go"]
