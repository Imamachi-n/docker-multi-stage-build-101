build:
	go build -o "bin/goServer"

run: build
	bin/goServer

dev:
	go run main.go

clean:
	rm -rf ./bin
