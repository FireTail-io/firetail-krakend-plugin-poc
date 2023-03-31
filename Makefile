# Params with defaults
ARCH := amd64
KRAKEND_BUILDER_IMAGE := builder:latest
RELEASE_VERSION := latest

.PHONY: build
build:
	docker run --platform linux/${ARCH} -v "${PWD}:/app" -w /app krakend/${KRAKEND_BUILDER_IMAGE} go build -buildmode=plugin -o firetail-krakend-plugin-${ARCH}-$(subst :,-,${KRAKEND_BUILDER_IMAGE})-${RELEASE_VERSION}.so .

.PHONY: test
test:
	go test ./... -race -coverprofile coverage.out -covermode atomic