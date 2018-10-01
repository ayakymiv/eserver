FROM golang:1.11-alpine3.8

RUN mkdir /app
ADD event /app/event
WORKDIR /app/event
RUN go build -o main .


