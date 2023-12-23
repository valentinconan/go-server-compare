FROM alpine:3.19.0
ARG VERSION
ENV APPLICATION_VERSION=$VERSION

WORKDIR /

COPY bin/go-api-amd64-linux /go-api-amd64-linux

EXPOSE 8080

ENTRYPOINT ["/go-api-amd64-linux"]
