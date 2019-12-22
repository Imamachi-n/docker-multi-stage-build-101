FROM golang:1.13.4

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /docker-multi-stage-build-101
COPY . .
RUN make build
EXPOSE 9000
ENTRYPOINT ["/docker-multi-stage-build-101/bin/goServer"]
