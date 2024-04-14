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

# migration:
# 	@read -p "Enter migration name: " name; \
# 		migrate create -ext sql -dir postgres/migrations $$name

migration:
	migrate create -dir postgres/migrations create_users_table


run:
	go run cmd/graphqlserver/*.go

generate: 
	go generate ./..