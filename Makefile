.PHONY: api
api:
	air

.PHONY: dev
dev:
	air

.PHONY: lint
lint:
	golangci-lint run