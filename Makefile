all: fmt lint

fmt:
	@printf "[CMD] "
	goimports -w -l .

lint:
	@printf "[CMD] "
	golint ./...

test:
	@printf "[CMD] "
	go test ./...

.PHONY: fmt lint all test
