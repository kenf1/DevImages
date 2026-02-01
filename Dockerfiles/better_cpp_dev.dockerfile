FROM emscripten/emsdk:5.0.0

RUN apt-get update && apt-get install -y \
    build-essential curl git neovim clang-format \
    software-properties-common

RUN add-apt-repository ppa:xmake-io/xmake && apt-get update && \
    apt-get install -y xmake