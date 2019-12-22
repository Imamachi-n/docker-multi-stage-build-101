# Builder image (intermediate container)
FROM golang:1.13.4 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /docker-multi-stage-build-101
COPY . .
RUN make build

# Runtime image
FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /docker-multi-stage-build-101/bin/goServer /goServer
EXPOSE 9000
ENTRYPOINT ["/goServer"]
