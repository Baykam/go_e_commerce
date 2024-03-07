.PHONY: database_start database_stop database_create

database_start:
	sudo service postgresql start

database_stop:
    sudo service postgresql stop

database_create:
    sudo service postgresql start
    sudo -u postgres psql
    ALTER USER postgres PASSWORD 'abc123';
    CREATE DATABASE micro_anthony_gg;