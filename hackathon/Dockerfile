FROM golang:1.19.2

WORKDIR /app

# Download necessary Go modules
ENV GO111MODULE=on
COPY go.mod ./
COPY go.sum ./

COPY main.go ./
COPY controller ./controller
COPY usecase ./usecase
COPY dao ./dao
COPY model ./model

RUN go mod download

CMD ["go", "run", "."]
