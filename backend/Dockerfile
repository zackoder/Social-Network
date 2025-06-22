FROM golang:latest

WORKDIR /backend
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server .

EXPOSE 8080
CMD ["./server"]
