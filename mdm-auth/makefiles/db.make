PROJECT_DIR = $(shell pwd)
PROJECT_BIN = $(PROJECT_DIR)/bin
MIGRATION_NAME ?= migration
MIGRATE_VERSION ?= v4.15.2
MIGRATE_ARGS ?= up
DB_PORT ?= 5432
DB_USER ?= postgres
DB_NAME ?= go-template
DB_DRIVER ?= postgres
DB_URL ?= "$(DB_DRIVER)://$(DB_USER)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable"

GOLANG_MIGRATE = $(PROJECT_BIN)/migrate

.PHONY: .install-migrate
.install-migrate:
	GOBIN=$(PROJECT_BIN) go install -tags '$(DB_DRIVER)' -mod=mod github.com/golang-migrate/migrate/v4/cmd/migrate@$(MIGRATE_VERSION)

migration: .install-migrate
	$(GOLANG_MIGRATE) create -ext sql -dir migrations $(MIGRATION_NAME)

migrate-up: .install-migrate
	$(GOLANG_MIGRATE) -path ./migrations -verbose -database 'postgres://mdm_user:fB2VPvcVtDc5lZhP3PkluYkh@10.10.99.1:5432/mdm?sslmode=disable&search_path=mdm_auth' up 1

migrate-down: .install-migrate
	$(GOLANG_MIGRATE) -path ./migrations -verbose -database 'postgres://mdm_user:fB2VPvcVtDc5lZhP3PkluYkh@10.10.99.1:5432/mdm?sslmode=disable&search_path=mdm_auth' down 1

migrate-force: .install-migrate
	$(GOLANG_MIGRATE) -path ./migrations -verbose -database 'postgres://mdm_user:fB2VPvcVtDc5lZhP3PkluYkh@10.10.99.1:5432/mdm?sslmode=disable&search_path=mdm_auth' force $(FORCE_VERSION)