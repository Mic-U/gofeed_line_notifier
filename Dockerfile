FROM golang:1.14.0 AS build-env

COPY . /go/src/nitifier
WORKDIR /go/src/nitifier
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep ensure
RUN go build -a -tags netgo -installsuffix netgo -ldflags="-s -w -extldflags \"-static\"" -o /go/src/nitifier/exec /go/src/nitifier/src

FROM busybox
COPY --from=build-env /go/src/nitifier/exec /usr/local/bin/notifier
CMD ["/usr/local/bin/notifier"]