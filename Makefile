GOPATH?=$(shell go env GOPATH)
TEST_PKG?=./...

.PHONY: generate-ssz
generate-ssz:
	@go generate ./...
