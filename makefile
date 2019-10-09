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

run:
	go run cmd/server/server.go