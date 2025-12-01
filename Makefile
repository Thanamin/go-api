# Makefile for convenient commands similar to `npm start`

.PHONY: start run build dev dev-run

start: ## Run the app (reads .env)
	set -a && [ -f .env ] && . ./.env || true && set +a && go run ./cmd/api

run: ## Alias for start
	$(MAKE) start

build: ## Build binary to ./bin/api
	go build -o bin/api ./cmd/api

# New target: build & run (used by dev mode)
dev-run:
	set -a && [ -f .env ] && . ./.env || true && set +a && \
	go build -o bin/api ./cmd/api && \
	./bin/api
# Dev mode with auto-reload
dev:
	$(shell go env GOPATH)/bin/reflex -s -r '\.go$$' -- sh -c '$(MAKE) dev-run'