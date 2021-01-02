FROM golang:1.15.6 AS build-env

COPY . /go/src/nitifier
WORKDIR /go/src/nitifier
RUN make build

FROM busybox
COPY --from=build-env /go/src/nitifier/dist/exec /usr/local/bin/notifier
COPY --from=build-env  /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/usr/local/bin/notifier"]