BIN_NAME = ard2rss
VERSION = 0.2.2

.PHONY: build

_bin:
	mkdir -p ./bin

build: _bin
	go build -o ./bin/${BIN_NAME} .

run: build
	./bin/${BIN_NAME}

docker:
	docker build -t mauamy/ard2rss:${VERSION} .

docker_push:
	docker push mauamy/ard2rss:${VERSION}

test: test_show test_collection

test_show:
	curl http://localhost:8080/sendung/levels-und-soundtracks/12642475

test_collection:
	curl -L http://localhost:8080/sammlung/einschlafen-mit-krimi-klassikern/58374573

clean:
	rm -rf ./bin