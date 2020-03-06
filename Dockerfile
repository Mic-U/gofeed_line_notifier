FROM golang:1.14.0 AS build-env

COPY . /go/src/nitifier
WORKDIR /go/src/nitifier
RUN make build

FROM busybox
COPY --from=build-env /go/src/nitifier/dist/exec /usr/local/bin/notifier
CMD ["/usr/local/bin/notifier"]