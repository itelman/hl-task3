FROM golang:1.22 AS builder

WORKDIR /
COPY . .

RUN go mod download

EXPOSE 8080

CMD ["go", "run", "."]
