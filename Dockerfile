FROM golang:1.20.6 as builder
ENV GOOS linux
ENV CGO_ENABLED 0
WORKDIR /app
COPY go.mod go.sum gen/ handler/ router/ ./
RUN go mod tidy
RUN go mod download
COPY *.go ./

RUN go build -o gRPC-Learning

# EXPOSE 3000
CMD ./gRPC-Learning