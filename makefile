gqlgen:
	go run github.com/99designs/gqlgen -v

generate:
	go generate ./...

run:
	go run -mod=vendor main.go