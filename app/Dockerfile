
ARG GO_VERSION=1.16

FROM golang:${GO_VERSION}-alpine AS build

RUN apk add --no-cache git


WORKDIR /app

# This container exposes port 8080 to the outside world
EXPOSE 8080

ENTRYPOINT ["tail", "-f", "/dev/null"]
