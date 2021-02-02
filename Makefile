# Binary name
BINARY=vault-creator

# Build values
VERSION=`git describe --tags --always`
BUILD=`date +%FT%T%z`

# Setup the -ldflags option for go build here
LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD}"

tests:
	# go test -v -short -race ./... # Disabled for now
	go vet ./...

build:
	@echo "Version: ${VERSION}"
	@echo "Build: ${BUILD}"
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build ${LDFLAGS} -installsuffix cgo -o ${BINARY}_v${VERSION} ./cmd/vault-creator/

run:
	go run ${LDFLAGS} ./cmd/vault-creator/main.go

clean:
	if [ -f ${BINARY}_v* ] ; then rm -f ${BINARY}_v* ; fi

.PHONY: all