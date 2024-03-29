FROM golang:1.22.1-alpine

WORKDIR /src

COPY . ./

RUN go build -o /bin/test

EXPOSE 8080
ENTRYPOINT /bin/test


