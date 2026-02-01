FROM golang:1.25.6-alpine3.22
RUN go telemetry off

RUN apk add --no-cache curl git make
RUN go install github.com/air-verse/air@latest
