# GOLANG SERVER
# Specify a base image
#FROM golang:1.18-alpine
FROM golang:1.18-buster

WORKDIR /go-server

# Install dependencies
# Playwright
RUN apt-get update
RUN apt-get -y install nodejs

RUN apt-get -y install libnss3
RUN apt-get -y install libnspr4
RUN apt-get -y install libatk1.0-0
RUN apt-get -y install ibatk-bridge2.0-0
RUN apt-get -y install libcups2
RUN apt-get -y install libdrm2
RUN apt-get -y install libdbus-1-3
RUN apt-get -y install libxcb1
RUN apt-get -y install libxkbcommon0
RUN apt-get -y install libx11-6
RUN apt-get -y install libxcomposite1
RUN apt-get -y install libxdamage1
RUN apt-get -y install libxext6
RUN apt-get -y install libxfixes3
RUN apt-get -y install libxrandr2
RUN apt-get -y install libgbm1
RUN apt-get -y install libpango-1.0-0
RUN apt-get -y install libcairo2
RUN apt-get -y install libasound2
RUN apt-get -y install libatspi2.0-0
RUN apt-get -y install libxshmfence1

# Go
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./cmd/api/*.go ./cmd/api/
COPY ./models/*.go ./models/

RUN go build -o /laundry-golang ./cmd/api

# Default command
CMD [ "/laundry-golang" ]