FROM golang:1.22.3
WORKDIR /backend
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server .

EXPOSE 8080
CMD ["./server"]