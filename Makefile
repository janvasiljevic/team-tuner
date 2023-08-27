yarn:
	cd frontend && \
		yarn

generate-api-bindings:
	cd backend && \
	swag fmt && \
	swag init && \
	cd ../frontend && \
	yarn run orval

admin: 
	cd backend && \
	go run scripts/add-user/main.go
	
db-migrate:
	cd backend && \
	atlas migrate apply \
	  --dir "file://resources/migrations" \
	  --url "postgres://dev:12345678@localhost:5432/dev?search_path=public&sslmode=disable"
	
seed-students:
	cd backend && \
	go run scripts/seed/seed.go