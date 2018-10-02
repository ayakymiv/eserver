FROM golang:1.11-alpine3.8
RUN apk update && apk add git gcc g++
RUN mkdir /go/src/event
ADD ./ /go/src/event
WORKDIR /go/src/event
RUN go get -d -v
RUN go install -v

FROM golang:1.11-alpine3.8
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/bin/event .
CMD ["./event"]


