# define PATH for db migration
export MIGRATION_DIR_PATH := db/migrations

LOCAL_BIN:=$(CURDIR)/bin-deps
PATH:=$(LOCAL_BIN):$(PATH)

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: clean
clean: ### Clean the build directory
	rm -rf ./bin/*

.PHONY: build
build: clean ### Build the binary file
	go build -o bin/blog cmd/main/main.go

.PHONY: debug
debug: ### Debug the main app, you need to attach the client to port 2345
	dlv debug --headless --only-same-user --listen :2345 --api-version 2 ./cmd/main/main.go -- -debug=true

.PHONY: watch
watch: ### Run make build to build the binary file and run it, will restart on file change
	air

.PHONY: gen-ent
gen-ent: ### Generate ent
	go generate ./ent

.PHONY: gen-grpc
gen-grpc: ### Generate grpc
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/src/*.proto

.PHONY: linter-golangci
linter-golangci: ### check by golangci linter
	golangci-lint run

.PHONY: gen-swagger
gen-swagger: ### Generate swagger documentation
	swag init -dir ./cmd/main/,./api -parseDependency -parseInternal

.PHONE: gen-keys
gen-keys: ### Generate private key and public key
	 go run -tags paseto_gen_key github.com/stewie1520/blog/cmd/gen_keys

bin-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/go-delve/delve/cmd/dlv@latest
	curl -sSf https://atlasgo.sh | sh -s -- -b $(LOCAL_BIN)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(LOCAL_BIN) v1.54.2
	GOBIN=$(LOCAL_BIN) go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(LOCAL_BIN)
	GOBIN=$(LOCAL_BIN) go install github.com/swaggo/swag/cmd/swag@latest
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
