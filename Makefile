APPLICATION := $(lastword $(subst /, ,$(dir $(CURDIR))))
PACKAGE := $(shell go list)/...
TESTS := $(wildcard *_test.go **/*_test.go)
SRC := $(filter-out $(TESTS), $(wildcard *.go **/*.go))
BIN := $(value GOPATH)\bin\$(APPLICATION).exe

# build when changed
$(BIN): $(SRC)
	go build -o $(BIN)

# run command at fixed intervals
watch: install
	@kouhai -i 2s 'make test'

# run all tests and rebuild when changed
test: $(BIN)
	@go test $(PACKAGE)

# build and install application
install: $(BIN)

# remove application
clean: 
	@go clean -i

.PHONY: clean test watch install
