export GO111MODULE = on


get-linter: 
	go get github.com/golangci/golangci-lint/cmd/golangci-lint

ci-lint:
	golangci-lint run
	go vet -composites=false -tests=false ./...
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -d -s
	go mod verify

install: 
	go install ./cmd/emd
	go install ./cmd/emcli

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download