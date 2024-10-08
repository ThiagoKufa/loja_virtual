FROM golang:1.23.1-alpine AS builder

ENV GOPATH=/go
ENV GOROOT=/usr/local/go
ENV PACKAGE=github.com/thiagokufa/loja_virtual
ENV BUILD_DIR=$GOPATH/src/$PACKAGE

COPY . ${BUILD_DIR}
WORKDIR ${BUILD_DIR}

RUN apk --no-cache add build-base bash && make clean build
RUN cp -v app /usr/bin/app

FROM alpine:latest
COPY --from=builder /usr/bin/app /usr/bin/app
COPY /entrypoint.sh /entrypoint.sh

RUN chmod 755 /entrypoint.sh

EXPOSE 3030
ENTRYPOINT ["/entrypoint.sh"]

