build:
	go build -o "bin/goServer"

run: build
	bin/goServer

dev:
	go run main.go

clean:
	rm -rf ./bin

dbuild1:
	docker image build . -f docker/01_raw/Dockerfile -t go-server-raw:dev

drun1:
	docker run --rm -p 9000:9000 go-server-raw:dev

dbuild2:
	docker image build . -f docker/02_multi-stage-build/Dockerfile -t go-server-multi:dev

drun2:
	docker run --rm -p 9000:9000 go-server-multi:dev
