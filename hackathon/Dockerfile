FROM golang:1.19.2

WORKDIR /app

# Download necessary Go modules
ENV GO111MODULE=on
COPY go.mod ./
COPY go.sum ./

COPY ./ .

RUN go mod download

CMD ["go", "run", "main.go"]
