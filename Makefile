# start server app hot reload
start:
	air

# Migrate scheme in ent to database
migrate_schema:
	go run ./cmd/migration/main.go

.PHONY: start migrate_schema