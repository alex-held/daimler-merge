.PHONY: gen lint test build install man clean

VERSION=v0.0.4
COMMIT  := `git rev-parse HEAD`
BINARY_NAME=daimler-merge
OUTPUT_DIR=out
BINARY_PATH=${OUTPUT_DIR}/${BINARY_NAME}


deps:
	go mod tidy
	go get -t -v
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.40.1

gen: deps
	go generate -tags=tools ./...

gen-ginkgo:
	go run github.com/onsi/ginkgo/ginkgo bootstrap
	go run github.com/onsi/ginkgo/ginkgo generate

lint: gen
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run --sort-results --fix --config .golangci.yaml  --print-resources-usage  --color auto --tests  --issues-exit-code 1 --uniq-by-line ./...

test: lint
	go run github.com/onsi/ginkgo/ginkgo -race -outputdir out -v  -coverprofile coverage.out -r  .  -- -test.v

test-watch:
	 go run github.com/onsi/ginkgo/ginkgo  -race -outputdir out -v -coverprofile coverage.out -r  .  -- -test.v

build:
	go build -o ${BINARY_PATH} -a -ldflags "-X=main.version=$(VERSION) -X=main.commit=$(COMMIT)" main.go

build-any:
	GOOS=darwin GOARCH=amd64 go build -o "${BINARY_PATH}-darwin-amd64" main.go
	GOOS=freebsd GOARCH=amd64 go build -o "${BINARY_PATH}-freebsd-amd64" main.go
	GOOS=linux GOARCH=amd64 go build -o "${BINARY_PATH}-linux-amd64" main.go
	GOOS=windows GOARCH=amd64 go build -o "${BINARY_PATH}-windows-amd64" main.go

clean:
	go clean
	rm ${OUTPUT_DIR}/*

install: build
	go install -a -ldflags "-X=main.version=$(VERSION) -X=main.commit=$(COMMIT)" ./...

man: gen
	go run main.go --help-man > daimler-merge.1

godoc: gen
	godoc -play -v  -links  -notes -index -http "localhost:6060"
