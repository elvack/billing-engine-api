run:
	go run main.go

run_db_migrate_up:
	go run database/migrate/migrate.go -t up

run_db_migrate_down:
	go run database/migrate/migrate.go -t down
