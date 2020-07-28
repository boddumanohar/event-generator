FROM golang:1.14.6 as build

ADD ./ /app

WORKDIR /app

RUN go build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app .
CMD ["./event-generator"]
