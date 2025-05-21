# Build
.PHONY: build
build:
	@go build -o packsmath cmd/api/main.go

# Container Build
.PHONY: docker-build
docker-build:
	@docker build -t packsmath .

.PHONY: docker-run
docker-run: docker-build
	@docker run -p 3000:3000 --rm packsmath

# Tests
.PHONY: tests
tests:
	@go clean -testcache
	@go test ./...

# Development
.PHONY: generate-mocks
generate-mocks:
	@[ -x "$$(go env GOPATH)/bin/moq" ] || go install github.com/matryer/moq@latest
	@go generate ./...

.PHONY: generate-swagger
generate-swagger:
	@[ -x "$$(go env GOPATH)/bin/swag" ] || go install github.com/swaggo/swag/cmd/swag@latest
	@$$(go env GOPATH)/bin/swag init --dir internal/handler/api --output internal/handler/api/docs --generalInfo api.go
