gqlgen:
	go run github.com/99designs/gqlgen -v

generate:
	go generate ./...

docker:
	docker build -t expenses:01 .
	docker run --name expenses -d expenses:01

rm:
	docker stop expenses
	docker rm expenses

run: godeps
	go run cmd/server/server.go

godeps:
ifeq (, $(shell which gqlgen))
	go get github.com/99designs/gqlgen
endif

watch: godeps
ifeq (, $(shell which wtc))
	curl -sfL --silent https://github.com/rafaelsq/wtc/releases/latest/download/wtc.linux64.tar.gz | tar -xzv && mv wtc $(go env GOPATH)/bin/
endif
	@wtc
