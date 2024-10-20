ROOT=go

.PHONY: run
run:
	go run http-servers.go

.PHONY: build
build:
	go build http-servers.go

.PHONY: ab
ab:
	@ab -n 10000 -c 100 http://localhost:8090/hello

