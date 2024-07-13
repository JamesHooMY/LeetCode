lint:
	golangci-lint run --timeout 10m ./... --fix

tidy:
	go mod tidy && go mod vendor

test:
	go clean -testcache && go test ./...

bench:
	go clean -testcache && go test -bench=. ./...

cover:
	go clean -testcache && go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out