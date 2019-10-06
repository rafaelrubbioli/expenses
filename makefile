gqlgen:
	go run github.com/99designs/gqlgen -v

generate:
	go generate ./...

docker:
	docker build -t expenses:01 .
	docker run -d expenses:01

rm:
	docker delete expenses:01

run:
	go run -mod=vendor main.go