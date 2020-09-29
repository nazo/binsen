.PHONY: migrate
migrate:
	docker-compose exec server goose -dir ./db/migrate postgres "host=postgres port=5432 user=binsen dbname=binsen_dev password=binsen sslmode=disable" up

.PHONY: sqlboiler
sqlboiler:
	docker-compose exec server sqlboiler --no-tests --wipe -o app/orm -p orm psql
	rm app/orm/goose_db_version.go

.PHONY: test
test:
	docker-compose exec server go test -v ./...

.PHONY: tools-install
tools-install:
	go install github.com/volatiletech/sqlboiler
	go install gopkg.in/godo.v2/cmd/godo
	go install github.com/pressly/goose/cmd/goose
	go install github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql
