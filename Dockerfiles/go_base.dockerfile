FROM golang:1.25.4-bookworm
RUN go telemetry off

RUN apt-get update && \
    apt-get install -y curl git make && \
    rm -rf /var/lib/apt/lists/*
RUN go install github.com/air-verse/air@latest