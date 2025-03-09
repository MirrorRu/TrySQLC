### GOPROXY='https://proxy.golang.org,direct'

### Work/Common ###
LOCAL_BIN:=$(CURDIR)/bin

#starting stub
.PHONY: starter-hello-stub
starter-hello-stub:
	$(info Hello!)
	$(info Firstly call "make setup"!)
	$(info other commands is listed above)
	grep "^[^\.]\w.*\:" makefile

PG_DSN="postgres://postgres:postgres@localhost:5432/tryall?sslmode=disable"
PG_SQL=migrations
.PHONY: pg-up
pg-up:
	goose postgres $(PG_DSN) -dir $(PG_SQL) up

.PHONY: pg-down
pg-down:
	goose postgres $(PG_DSN) -dir $(PG_SQL) down

.PHONY: .cls
.cls:
	clear

.PHONY: setup-tools
setup-tools:
	go get github.com/dmarkham/enumer
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

.PHONY: generate
generate: gen

.PHONY: gen
gen:
	$(info code generating )
	go generate ./...

