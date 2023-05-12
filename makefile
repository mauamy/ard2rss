BIN_NAME = ard2rss

.PHONY: build

_bin:
	mkdir -p ./bin

build: _bin
	go build -o ./bin/${BIN_NAME} .

run: build
	./bin/${BIN_NAME}

clean:
	rm -rf ./bin