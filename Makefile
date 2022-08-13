run :
	go run ./cmd/api

psql:
	psql ${GREENLIGHT_DB_DSN}

up :
	@echo 'Running migrations....'
	migrate -path ./migrations -database ${GREENLIGHT_DB_DSN}  up