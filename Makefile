postgres:
	sudo docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	sudo docker exec -it postgres12 createdb --username=root --owner=root bank_trans

dropdb: 
	sudo docker exec -it postgres12 dropdb bank_trans

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/bank_trans?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/bank_trans?sslmode=disable" -verbose down

sqlc:
	sudo docker run --rm -v ${CURDIR}/db:/src -w /src kjconroy/sqlc:1.14.0 generate

test:
	go test -v -cover ./...

.PHONY:postgres createdb dropdb migrateup migratedown sqlc