SHELL = /bin/bash
ROOT_DIR = $(shell pwd)

GO ?= go
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")
GOBIN = $(ROOT_DIR)/bin
export PATH := $(GOBIN):$(PATH)

TEST_DIRS := $(shell $(GO) list ./... | grep -v mocks)

.PHONY: test
test:
	$(GO) test -cover $(TEST_DIRS)

.PHONY: fmt
# Ensure consistent code formatting.
fmt:
	$(GOFMT) -w $(GOFILES)

.PHONY: fmt-check
# format (check only).
fmt-check:
	@diff=$$($(GOFMT) -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

.PHONY: coverage
coverage:
	# lcov
	$(GO) test -cover $(TEST_DIRS) -coverprofile=coverage.out
	bin/gcov2lcov -infile=coverage.out -outfile=coverage.lcov
	genhtml coverage.lcov -o outputs
	# go coverage
	$(GO) tool cover -html=coverage.out -o outputs/gocoverage.html

.PHONY: tools
tools:
	GOBIN=$(GOBIN) $(GO) install github.com/jandelgado/gcov2lcov@v1.0.5
