# Build files
FROM golang:alpine as BUILD

COPY . .

RUN apk add --no-cache $PACKAGES && \
  make install

## Copy from build only cli
FROM alpine

WORKDIR /root

COPY --from=BUILD /go/bin/emd /usr/bin/emd
COPY --from=BUILD /go/bin/emcli /usr/bin/emcli

CMD emd