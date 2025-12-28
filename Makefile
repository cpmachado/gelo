MAIN=.
TAG_VERSION=$(shell git describe --tags --exact-match 2>/dev/null || git rev-parse --short HEAD)
VERSION=$(subst /,_,$(TAG_VERSION))

all: build
	@echo all built

build:
	 go build $(MAIN)

clean:
	@rm -rf gelo target
	@echo all removed

lint:
	golangci-lint run ./...
	@echo all code is linted

format:
	gofmt -w -s .

format-check:
	gofmt -l .

run:
	go run $(MAIN) -loglevel DEBUG

test:
	go test -v ./...

sbom: build
	@mkdir -p target/sbom
	cyclonedx-gomod bin -json -output ./target/sbom/gelo-$(VERSION).bom.json ./gelo

setup-devtools:
	# Vulnerability checker
	go install golang.org/x/vuln/cmd/govulncheck@latest
	# Debugger
	go install github.com/go-delve/delve/cmd/dlv@latest
	# Linter
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	# SBOM generator
	go install github.com/CycloneDX/cyclonedx-gomod/cmd/cyclonedx-gomod@latest


.PHONY: all build clean dev-dependencies lint format format-check run test sbom setup-devtools
