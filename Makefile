postgres:
	sudo docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	sudo docker exec -it postgres12 createdb --username=root --owner=root bank_trans

dropdb: 
	sudo docker exec -it postgres12 dropdb bank_trans

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/bank_trans?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/bank_trans?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/bank_trans?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/bank_trans?sslmode=disable" -verbose down 1


sqlc:
	sudo docker run --rm -v ${CURDIR}/db:/src -w /src kjconroy/sqlc:1.14.0 generate

test:
	go test -v -cover ./...
server:
	go run main.go

#For mockgen package. There is the first path is new generated mock model, second one is current real model
gomock:
	mockgen -package mockdb -destination db/mock/store.go bank/db/sqlc Store


.PHONY:postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server gomock