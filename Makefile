run:
	go run main.go

run_docker:
	docker stop billing-engine-api || true
	docker rm billing-engine-api || true
	docker build --progress=plain -t billing-engine-api:develop .
	docker run -d --name billing-engine-api -p 8000:8000 billing-engine-api:develop

run_db_migrate_up:
	go run database/migrate/migrate.go -t up

run_db_migrate_down:
	go run database/migrate/migrate.go -t down

run_db_seed_admin:
	go run database/seed/admin/admin.go
