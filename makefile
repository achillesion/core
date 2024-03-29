export GO111MODULE = on


get-linter: 
	go get github.com/golangci/golangci-lint/cmd/golangci-lint

lint:
	golangci-lint run
	go vet -composites=false -tests=false ./...
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -d -s
	go mod verify

all: install
install: go.sum
		GO111MODULE=on go install -mod=readonly -tags "$(build_tags)" ./cmd/dtd
		GO111MODULE=on go install -mod=readonly -tags "$(build_tags)" ./cmd/dtcli

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify