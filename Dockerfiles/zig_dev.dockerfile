FROM fedora:44

WORKDIR /app

RUN dnf update -y && dnf install -y coreutils \
    make gcc glibc-devel git zig