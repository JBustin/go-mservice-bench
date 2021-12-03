FROM golang:1.17.3
RUN apt-get update && apt-get install -y vim curl
RUN mkdir -p /usr/app
WORKDIR /usr/app
COPY . .