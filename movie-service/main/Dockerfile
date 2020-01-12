FROM golang:latest as builder

WORKDIR /movie-service

COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy pre-built binary file from previous stage
COPY --from=builder /movie-service/server .
COPY --from=builder /movie-service/.env .

EXPOSE 8081

CMD ["./server"]