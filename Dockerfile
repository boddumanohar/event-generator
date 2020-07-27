FROM golang:1.14.6 as build

ADD ./ /app

WORKDIR /app

RUN go build

# use the same docker image for both build and publishing
# we should copy the image and distribute only the binary.
# so the image used for publishing should be extremely light weight.
# lighter than the golang:1.14.6 itself.
FROM build

