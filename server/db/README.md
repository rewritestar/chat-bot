## migrate 사용법

$ migrate create -ext sql -dir migrations -seq create_table_user 1

$ migrate -source file://db/migrations -database "mysql://root:root@tcp(localhost:3306)/chatbot" up

$ migrate -source file://db/migrations -database "mysql://root:root@tcp(localhost:3306)/chatbot" version

drop table schema_migrations;

https://github.com/golang-migrate/migrate/tree/master/database/mysql

https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
