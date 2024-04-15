mock:
	mockery --all --keeptree

migrate:
	migrate -source file://postgres/migrations \
			-database postgres://xuser:xpassword@127.0.0.1:5432/x_clone_development?sslmode=disable up

rollback:
	migrate -source file://postgres/migrations \
			-database postgres://xuser:xpassword@127.0.0.1:5432/x_clone_development?sslmode=disable down 1

drop:
	migrate -source file://postgres/migrations \
			-database postgres://xuser:xpassword@127.0.0.1:5432/x_clone_development?sslmode=disable drop

migration:
	@read -p "Enter migration name: " name; \
		migrate create -ext sql -dir postgres/migrations $$name

# migration:
# 	migrate create -ext sql -dir postgres/migrations/ -seq -digits 6 create_users_table

run:
	lsof -i :8080 | awk 'NR!=1 {print $$2}' | xargs -r kill -9
	go run cmd/graphqlserver/*.go

generate: 
	go generate ./...

history:
	cat ~/.zsh_history | grep "make" | sort | uniq -c | sort -nr