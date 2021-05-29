.PHONY: gen lint test build install man clean

VERSION=v0.0.2
COMMIT  := `git rev-parse HEAD`
BINARY_NAME=daimler-merge
OUTPUT_DIR=out
BINARY_PATH=${OUTPUT_DIR}/${BINARY_NAME}

deps:
	go mod tidy
	go get -t -v

gen: deps
	go generate -tags=tools ./...

gen-ginkgo:
	go run github.com/onsi/ginkgo/ginkgo bootstrap
	go run github.com/onsi/ginkgo/ginkgo generate

lint: gen
	 golangci-lint run --sort-results --fix --config .golangci.yaml  --print-resources-usage  --color auto --tests  --issues-exit-code 1 --uniq-by-line ./...

test: lint
	go run github.com/onsi/ginkgo/ginkgo -race -r  -v -cover -coverprofile=out/coverage.out -r  .  -- -test.v

test-watch:
	 go run github.com/onsi/ginkgo/ginkgo watch -race -r  -v -cover -coverprofile=out/coverage.out -r  .  -- -test.v

build: test
	go build -o ${BINARY_PATH} -a -ldflags "-X=main.version=$(VERSION) -X=main.commit=$(COMMIT)" main.go

clean:
	go clean
	rm ${BINARY_PATH}/*

install: build
	go install -a -ldflags "-X=main.version=$(VERSION) -X=main.commit=$(COMMIT)" ./...

man: gen
	go run main.go --help-man > daimler-merge.1

godoc: gen
	godoc -play -v  -links  -notes -index -http "localhost:6060"
