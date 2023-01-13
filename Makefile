include .env

run:
	go run cmd/main.go

build:
	go build -o todo-rest-api cmd/main.go

rundb:
	docker run -d --name todo-db -e POSTGRES_PASSWORD=${DB_PASSWORD} --net=host --rm postgres:9.4

stopdb:
	docker stop todo-db

migrcreate:
	migrate create -ext sql -dir ./schema -seq init

migrup:
	migrate -path ./schema -database 'postgres://postgres:${DB_PASSWORD}@localhost:5432/postgres?sslmode=disable' up

migrdown:
	migrate -path ./schema -database 'postgres://postgres:${DB_PASSWORD}@localhost:5432/postgres?sslmode=disable' down

image: rmimage
	sudo docker build -t todo-rest-api .

rmimage:
	docker image rm -f todo-rest-api

cont:
	docker run -dit --rm --net=host --name todo-rest-api todo-rest-api

stopcont:
	docker stop todo-rest-api

prune:
	docker image prune