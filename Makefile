#@IgnoreInspection BashAddShebang
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export CGO_ENABLED=0
export GO111MODULE=on

.PHONY: help
help: ## Shows help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.DEFAULT_GOAL := help

.which-go:
	@which go > /dev/null || (echo "install go from https://golang.org/dl/" & exit 1)

.PHONY: format
format: .which-go ## Formats Go files
	gofmt -s -w $(ROOT)

.PHONY: lint
lint: .which-lint ## Checks code with Golang CI Lint
	golangci-lint run
