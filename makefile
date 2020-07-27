DIRS_WITH_TESTS := $(shell find * -type f -name '*_test.go' | xargs -n 1 dirname | sort | uniq)
DIRS_WITH_GOFILES := $(shell find * -type f -name '*.go' | xargs -n 1 dirname | sort | uniq)

myapp: go.mod *.go pkg/*/*/*.go makefile
	go build -o myapp

fmt:
	for dir in $(DIRS_WITH_GOFILES); do (cd $${dir}; go fmt); done

test:
	for dir in $(DIRS_WITH_TESTS); do (cd $${dir}; go test); done
