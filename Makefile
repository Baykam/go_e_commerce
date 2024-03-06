.PHONY: doc unittest mock database

doc:
	swag fmt && swag init -g ./cmd/api/main.go

unittest:
	go test -timeout 9000s -a -v -coverprofile=coverage.out -coverpkg=./... ./... 2>&1 | tee report.out

mock:
	go generate ./...

database:
	sudo service postgresql start
	# sudo -u postgres psql
	# ALTER USER postgres PASSWORD 'abc123';
	# CREATE DATABASE micro_anthony_gg;
