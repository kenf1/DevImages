FROM python:3.13.5-slim

RUN apt-get update && apt-get install -y \
    build-essential \
    curl \
    software-properties-common \
    git \
    && rm -rf /var/lib/apt/lists/*
RUN python3 -m pip install --upgrade pip

COPY ./PyDS/requirements.txt ./requirements.txt
RUN python3 -m pip install -r ./requirements.txt