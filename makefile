.PHONY: all
all: test

.PHONY: test
test:
	@echo "test"
	@go test -v ./...