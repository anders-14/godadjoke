all: fmt lint

fmt:
	@printf "[CMD] "
	goimports -w -l .

lint:
	@printf "[CMD] "
	golint ./...

.PHONY: all fmt lint
