# docker-multi-stage-build-101

## How to build each docker container image

Example:

```bash
docker image build . -f docker/02_multi-stage-build/Dockerfile -t go-server-raw:dev
docker run --rm -p 9000:9000 go-server-raw:dev
```
