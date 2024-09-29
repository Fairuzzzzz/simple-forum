export MYSQL_URL = 'mysql://root:testpassword@tcp(127.0.0.1:3306)/belajar'

migrate-create:
	@ migrate create -ext sql -dir scripts/migrations -seq $(name)

migrate-up:
	@ migrate -database ${MYSQL_URL} -path scripts/migrations up

migrate-down:
	@ migrate -database ${MYSQL_URL} -path scripts/migrations down
