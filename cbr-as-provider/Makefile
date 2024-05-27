.PHONY: all
all: format install_tools build
#static_analysis

.PHONY: format
format:
	@find . -type f -name "*.go*" -not -path "./vendor/*" -print0 | xargs -0 gofmt -s -w

# Performing linting and static analysis
.PHONY: lint
lint: static_analysis

# Performing linting and static analysis
.PHONY: static_analysis
static_analysis:
	GO111MODULE=on $(shell go env GOPATH)/bin/golangci-lint run --enable gofmt -E golint -E gosec --exclude-use-default=false

.PHONY: gosec
gosec:
	GO111MODULE=on $(shell go env GOPATH)/bin/gosec ./...

# Tools installation
.PHONY: install_tools
install_tools:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.30.0
	curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v2.5.0

.PHONY: clean
clean:
	@go clean ./...

.PHONY: build
build: build_service_zone_utils_cbr_as_provider

.PHONY: build_service_zone_utils_cbr_as_provider
build_service_zone_utils_cbr_as_provider:
	@mkdir -p ./bin/cbr-as-provider  
	@go build -o ./bin/cbr-as-provider ./cbr-as-provider/... 
