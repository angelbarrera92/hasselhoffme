.PHONY: test lint check cover build
#.PHONY: release

test: ## Run tests
	go test -race ./... -timeout=5m

lint: ## Run linters. Use make install-linters first.
	go get -u golang.org/x/lint/golint
	golint --set_exit_status ./...
	go vet -all ./...

check: lint test  ## Run tests and linters

cover: ## Runs tests on ./cmd/ with HTML code coverage
	go test -race -cover -coverprofile=cover.out -coverpkg=github.com/angelbarrera92/hasselhoffme/... ./...
	go tool cover -html=cover.out

build: ## Builds the binary
	export GO111MODULE=on && \
		go mod download && \
		go build -v ./cmd/hasselhoffme

## TODO: Add support for goreleaser
#release: check	## Use GoReleaser to build, package and release
#	goreleaser --rm-dist
