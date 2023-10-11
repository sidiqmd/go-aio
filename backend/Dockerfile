# syntax=docker/dockerfile:1

FROM golang:1.19

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/*.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-go-todo

EXPOSE 8080

# Run
CMD ["/docker-go-todo"]