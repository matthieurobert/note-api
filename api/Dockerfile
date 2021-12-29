FROM golang:alpine3.15 AS builder
WORKDIR /go/src/github.com/matthieurobert/amp/api
COPY . /go/src/github.com/matthieurobert/amp/api/
RUN set -xe; \
    apk add --no-cache git && \
    go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api-binary

#       FINAL IMAGE
FROM alpine:3.15
#   LABELLING
LABEL maintainer="Matthieu ROBERT <matthieurobert82@gmail.com>"
LABEL version="alpha-0.1"
LABEL description="A docker container to run the amp API"
#   SETUP
WORKDIR /app
COPY --from=builder /go/src/github.com/matthieurobert/amp/api .
EXPOSE 8000
#   RUN
CMD ./api-binary