.PHONY: test vendor

test:
	go clean -testcache
	go test -race -v ./...

vendor:
	go mod tidy
	go mod vendor
