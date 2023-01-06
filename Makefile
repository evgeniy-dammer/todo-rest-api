include .env

run:
	go run cmd/main.go

build:
	go build -o todo-rest-api cmd/main.go

rundb:
	docker run -d --name todo-db -e POSTGRES_PASSWORD=${DB_PASSWORD} -p 5432:5432 --rm postgres:9.4

stopdb:
	docker stop todo-db

migrcreate:
	migrate create -ext sql -dir ./schema -seq init

migrup:
	migrate -path ./schema -database 'postgres://postgres:${DB_PASSWORD}@localhost:5432/postgres?sslmode=disable' up

migrdown:
	migrate -path ./schema -database 'postgres://postgres:${DB_PASSWORD}@localhost:5432/postgres?sslmode=disable' down