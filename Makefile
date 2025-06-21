PHONY: migrate_up migrate_down

migrate_up:
	..\bin\goose.exe postgres "user=fedor password=fedor_pass dbname=rate port=15432 host=localhost sslmode=disable" up -dir .\migrations

migrate_down:
	..\bin\goose.exe postgres "user=fedor password=fedor_pass dbname=rate port=15432 host=localhost sslmode=disable" down -dir .\migrations