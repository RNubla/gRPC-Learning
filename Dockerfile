FROM golang:1.20.6 as builder
ENV GOOS linux
ENV CGO_ENABLED 0
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go mod download
COPY *.go ./

RUN go build -o gRPC-Learning

EXPOSE 6000
CMD ./gRPC-Learning