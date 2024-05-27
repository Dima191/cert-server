generate_code:
	protoc -I ./api/cert ./api/cert/cert.proto --go_opt=paths=source_relative --go_out=./pkg/cert --go-grpc_opt=paths=source_relative --go-grpc_out=./pkg/cert

export_env:
	export dbURL=driver://user:your_password@localhost:port/db_name && export migratePath=./migrations

migrate_up:
	go build ./cmd/migrator && ./migrator --mode=up && rm -rf migrator
migrate_down:
	go build ./cmd/migrator && ./migrator --mode=down && rm -rf migrator

run:
	go run ./cmd/cert --debug=true

.PHONY: generate_code, export_env, migrate_up, migrate_down, run

.DEFAULT_GOAL = run