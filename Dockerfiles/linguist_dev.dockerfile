FROM rust:1.91.1-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential libssl-dev pkg-config curl git && \
    rm -rf /var/lib/apt/lists/*

RUN cargo install repo-analyzer