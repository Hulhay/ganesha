run:
	go run api/main.go

generate:
	go run github.com/99designs/gqlgen generate

test:
	go test ./... -coverprofile=cover.out

cover: test
	go tool cover -html=cover.out -o cover.html; \
	open cover.html

migrate-up:
	migrate -path migrations -database "postgresql://[YOUR_DB_USERNAME]:[YOUR_DB_PASSWORD]@[YOUR_DB_HOST]:[YOUR_DB_PORT]/[YOUR_DB_NAME]" -verbose up

migrate-down:
	migrate -path migrations -database "postgresql://[YOUR_DB_USERNAME]:[YOUR_DB_PASSWORD]@[YOUR_DB_HOST]:[YOUR_DB_PORT]/[YOUR_DB_NAME]" -verbose down