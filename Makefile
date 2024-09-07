SHELL = /bin/bash
ROOT_DIR = $(shell pwd)

GO ?= go
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")
GOBIN = $(ROOT_DIR)/bin
export PATH := $(GOBIN):$(PATH)

.PHONY: test
test:
	$(GO) test -cover ./internal/...

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
	$(GO) test -cover ./... -coverprofile=coverage.out
	bin/gconv2lconv -infile=coverage.out -outfile=coverage.lcov
	genhtml coverage.lcov -o site
	# go coverage
	$(GO) tool cover -html=coverage.out -o site/gocoverage.html

.PHONY: tools
tools:
	GOBIN=$(GOBIN) $(GO) install github.com/jandelgado/gcov2lcov@v1.0.5
