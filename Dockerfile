FROM golang:1.19-alpine3.17 as bulder

RUN apk add --no-cache ca-certificates
RUN apk update --no-cache && apk upgrade --no-cache
RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories
RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories
RUN apk add libc-dev gcc openssh-client git --no-cache

WORKDIR follower-service

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app \
    -ldflags "-X main.revision=$(git rev-parse --abbrev-ref HEAD)-$(git describe --abbrev=7 --always --tags) -s -w" \
    -tags musl \
    ./cmd/follower

FROM alpine:3.17
RUN apk add --no-cache ca-certificates
RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories
RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories
RUN apk update && apk upgrade
COPY --from=builder /app /app
EXPOSE 3000
CMD ["/app"]