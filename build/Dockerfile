#
# STEP 1 - build
# 

FROM docker.io/library/golang:alpine as builder

# fresh certs
RUN apk update && apk add --no-cache git ca-certificates tzdata make && update-ca-certificates

WORKDIR ${GOPATH}/src/github.com/infinite-industries/infinite-bot
COPY . .

# Fetch dependencies.
RUN cd cmd/discord-bot && \
    go get -d -v

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo \
  -o /go/bin/discord-bot ./cmd/discord-bot

#
# Step 2 - small image
#

FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

# copy static executable from builder
COPY --from=builder /go/bin/discord-bot /app/discord-bot

WORKDIR /app

USER nobody

ENV BOT_TOKEN=""

# prometheus metrics exposition
EXPOSE 2112/tcp

ENTRYPOINT ["./discord-bot"]

