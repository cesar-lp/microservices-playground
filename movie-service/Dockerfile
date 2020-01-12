FROM golang:latest as builder

# Set current working directory
WORKDIR /movie-service

COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./main

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/
RUN mkdir ~/app

# Copy pre-built binary file from previous stage
COPY --from=builder /movie-service/main/main ./app
COPY --from=builder /movie-service/.env .

EXPOSE 8081

CMD ["./app/main"]