all: test check bench

deps:
	go get -v -d -t ./...
	go get -v github.com/alecthomas/gometalinter
	gometalinter --install

test: deps
	go test -race -v -coverprofile=coverage.txt -covermode=atomic

bench: deps
	go test -v -run ^Test$$ -bench=. ./... -gocheck.b

check:
	go test -i
	gometalinter --vendored-linters --deadline=30s --cyclo-over=16 ./...

.PHONY: deps test bench check
