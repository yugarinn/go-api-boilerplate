start_database:
	docker compose up --build

start_development_server:
	BOILERPLATE_ENV=development ${GOPATH}/bin/gin run main.go

migrate:
	${GOPATH}/bin/migrate -path ./database/migrations -database "mysql://boilerplate:secret@tcp(localhost:33060)/boilerplate?charset=utf8mb4&parseTime=True&loc=Local" up

test:
	BOILERPLATE_ENV=test ${GOPATH}/bin/gotestsum --format dots ./tests/feature
