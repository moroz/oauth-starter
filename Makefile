install:
	go get
	which modd || go install github.com/cortesi/modd/cmd/modd@latest

migrate:
	scripts/migrate_db.sh

rollback:
	scripts/migrate_db.sh down

