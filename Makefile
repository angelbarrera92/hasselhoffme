.PHONY: build

build: ## Builds the binary
	go mod download && go mod tidy && go build -v ./cmd/hasselhoffme
