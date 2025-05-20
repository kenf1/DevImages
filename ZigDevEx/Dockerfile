FROM ghcr.io/kenf1/zigdev:latest

WORKDIR /app
COPY . .

RUN dnf update -y
RUN zig fetch --save "git+https://github.com/zigzap/zap#v0.10.1"

CMD ["zig", "build", "run"]