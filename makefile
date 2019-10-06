MYSQL_PORT ?= 3306
DATABASE_NAME ?= mariadb
SENHA_ROOT ?= root

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
	docker run -d -p $(MYSQL_PORT) --name=$(DATABASE_NAME) -v ${PWD}/mysql-data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=$(SENHA_ROOT) mariadb
	go run -mod=vendor main.go

clean:
	docker rm -f $(DATABASE_NAME)