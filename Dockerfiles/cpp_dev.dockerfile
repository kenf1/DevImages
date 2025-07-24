FROM ubuntu:25.10

RUN apt-get update && apt-get install -y \
    build-essential curl software-properties-common \
    libboost-all-dev git