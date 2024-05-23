run:
	go run cmd/http/api.go

migrate-up:
	migrate -database "mysql://root:123456@tcp(localhost:3306)/user" -path db/migrations up

migrate-down:
	migrate -database "mysql://root:123456@tcp(localhost:3306)/user" -path db/migrations down
