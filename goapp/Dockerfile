FROM --platform=$BUILDPLATFORM golang:1.20.6-alpine as builder

WORKDIR /code

ENV CGO_ENABLED 0

ENV GOPATH /go

ENV GOCACHE /go-build

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod/cache \
    go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod/cache \
    --mount=type=cache,target=/go-build \
    go build -o bin/goapp main.go

CMD ["/code/bin/goapp"]

FROM builder AS dev-envs

RUN <<EOF
apk update
apk add git
EOF

RUN <<EOF
addgroup -S docker
adduser -S --shell /bin/bash --ingroup docker vscode
EOF

# install Docker tools (cli, buildx, compose)
COPY --from=gloursdocker/docker / /

CMD ["go", "run", "main.go"]

FROM scratch
COPY --from=builder /code/bin/goapp /usr/local/bin/goapp
CMD ["/usr/local/bin/goapp"]
# RUN go mod tidy

# RUN go mod download

# COPY *.go ./

# RUN go build -o goapp

# EXPOSE 70:70

# CMD ./goapp