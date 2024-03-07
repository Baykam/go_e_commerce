.PHONY: migrate migrate_create migrate_up

migrate:
    export POSTGRESQL_URL='postgres://postgres:abc123@localhost:5432/golang_test?sslmode=disable' \
    echo $POSTGRESQL_URL \
    migrate -database ${POSTGRESQL_URL} up

migrate_create:
    migrate create -ext sql -dir migrations -seq init_schema

migrate_up:
    export POSTGRESQL_URL='postgres://postgres:abc123@localhost:5432/golang_test?sslmode=disable'; \
    migrate -database ${POSTGRESQL_URL} up
