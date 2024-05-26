run:
	go run cmd/http/api.go

migrate-up:
	migrate -database "mysql://syamsul:SyamSul1234#@tcp(18.139.84.244:3306)/user" -path db/migrations up

migrate-down:
	migrate -database "mysql://root:123456@tcp(localhost:3306)/user" -path db/migrations down
