SWAG=$(go env GOPATH)/bin/swag

generate-mocks:
	@[ -x "$$(go env GOPATH)/bin/moq" ] || go install github.com/matryer/moq@latest
	@go generate ./...

generate-swagger:
	@[ -x "$$(go env GOPATH)/bin/swag" ] || go install github.com/swaggo/swag/cmd/swag@latest
	@$$(go env GOPATH)/bin/swag init --dir internal/handler/api --output internal/handler/api/docs --generalInfo api.go
