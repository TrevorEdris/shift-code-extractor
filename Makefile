GOOS := linux

all: binaries

.PHONY: bin
bin:
	mkdir bin || true

.PHONY: extractor
extractor: bin
	go build -o bin/extractor cmd/extractor/main.go

.PHONY: login
login: bin
	go build -o bin/login cmd/login/main.go

.PHONY: subscribe
subscribe: bin
	go build -o bin/subscribe cmd/subscribe/main.go

.PHONY: subscribeui
subscribeui: bin
	go build -o bin/subscribeui cmd/subscribe-ui/main.go

.PHONY: binaries
binaries: extractor login subscribe subscribeui

.PHONY: zip
zip:
	./create-zips
