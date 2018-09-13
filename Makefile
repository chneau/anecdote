.SILENT:
.ONESHELL:
.NOTPARALLEL:
.EXPORT_ALL_VARIABLES:
.PHONY: run deps build clean exec

run: build exec clean

exec:
	./bin/app

build:
	CGO_ENABLED=0 go build -o bin/app -ldflags '-s -w -extldflags "-static"'

clean:
	rm -rf bin
	rm -rf upload

deps:
	govendor init
	govendor add +e
	govendor update +v

dev:
	go get -u -v github.com/kardianos/govendor

