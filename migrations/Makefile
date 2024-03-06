migrate:
    export POSTGRESQL_URL='postgres://postgres:abc123@localhost:5432/micro_anthony_gg?sslmode=disable'
	echo $POSTGRESQL_URL
	migrate -database ${POSTGRESQL_URL} -path migrations up