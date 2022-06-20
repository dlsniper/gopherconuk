FROM golang:1.11beta2-alpine3.8 AS build-env

# Allow Go to retrive the dependencies for the build step
RUN apk add --no-cache git

# Secure against running as root
RUN adduser -D -u 10000 florin
RUN mkdir /go-server-template/ && chown florin /go-server-template/
USER florin

WORKDIR /go-server-template/
ADD . /go-server-template/

# Compile the binary, we don't want to run the cgo resolver
RUN CGO_ENABLED=0 go build -o /go-server-template/goServerT .

# final stage
FROM alpine:3.8

# Secure against running as root
RUN adduser -D -u 10000 florin
USER florin

WORKDIR /
COPY --from=build-env /go-server-template/certs/docker.localhost.* /
COPY --from=build-env /go-server-template/goServerT /

EXPOSE 8080

CMD ["/goServerT"]
