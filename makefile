myapp: go.mod *.go pkg/*/*/*.go makefile
	go build -o myapp

fmt:
	go fmt
	(cd pkg/github/release/; go fmt)
