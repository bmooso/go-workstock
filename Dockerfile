FROM golang:alpine
ARG pkg=github.com/bmooso/go-helloworld

COPY . $GOPATH/src/$pkg
WORKDIR $GOPATH/src/$pkg

RUN set -ex && \
  CGO_ENABLED=0 GOOS=linux \
  go build \
  -a \
  -installsuffix cgo \
  -o main

FROM alpine
ARG pkg=github.com/bmooso/go-helloworld
COPY --from=0 /etc/ssl /etc/ssl
COPY --from=0 /go/src/$pkg/main /app/main
EXPOSE 8080
WORKDIR /app
CMD ["/app/main", "--port=8080", "--host=0.0.0.0"]
