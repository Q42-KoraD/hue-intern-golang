FROM golang:1.22

# Set destination for COPY
WORKDIR /app

COPY . /app

RUN go mod download

# Build
RUN go build -o server main.go

EXPOSE 3000

# Run
CMD ["./server"]
