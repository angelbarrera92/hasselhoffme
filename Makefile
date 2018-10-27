.PHONY: test lint check cover
.PHONY: install-linters
#.PHONY: release

test: ## Run tests
	go test -race ./... -timeout=5m

lint: ## Run linters. Use make install-linters first.
	vendorcheck ./...
	$(GOPATH)/bin/golangci-lint run --no-config --deadline=3m --disable-all --tests \
		-E golint \
		-E goimports \
		-E varcheck \
		-E unparam \
		-E deadcode \
		-E structcheck \
		-E errcheck \
		-E gosimple \
		-E staticcheck \
		-E unused \
		-E ineffassign \
		-E typecheck \
		-E gas \
		-E megacheck \
		-E misspell \
		./...
	# The govet version in golangci-lint is out of date and has spurious warnings, run it separately
	#go vet -all ./...

check: lint test  ## Run tests and linters

cover: ## Runs tests on ./cmd/ with HTML code coverage
	go test -race -cover -coverprofile=cover.out -coverpkg=github.com/angelbarrera92/hasselhoffme/... ./...
	go tool cover -html=cover.out

build: ## Builds the binary
	export GO111MODULE=on
	go mod download
	go build -v

install-linters: ## Install linters
	go get -u github.com/FiloSottile/vendorcheck
	# Pin to v1.10.2
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b $(GOPATH)/bin v1.10.2

## TODO: Add support for goreleaser
#release: check	## Use GoReleaser to build, package and release
#	goreleaser --rm-dist
