postgres:
	docker run --name=postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it postgres12 createdb --username=root owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrate-up:
	migrate -path=db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrate-up1:
	migrate -path=db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migrate-down:
	 migrate -path=db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migrate-down1:
	migrate -path=db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

test:
	go test -v -cover ./...

sqlc:
	sqlc generate

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/hagios2/simple-bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrate-up migrate-down sqlc test server mock migrate-down1 migrate-up1
