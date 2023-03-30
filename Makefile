# Params with defaults
ARCH := amd64
KRAKEND_VERSION := 2.2.1
RELEASE_VERSION := latest

build:
	docker run --platform linux/${ARCH} -v "${PWD}:/app" -w /app krakend/builder:${KRAKEND_VERSION} go build -buildmode=plugin -o firetail-krakend-plugin-${ARCH}-${KRAKEND_VERSION}-${RELEASE_VERSION}.so .
