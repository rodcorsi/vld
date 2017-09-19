.PHONY: build \
test \
cover

GO := go

test:
	@echo Running tests
	$(GO) test -run=$(TESTS) -test.v -test.timeout=120s -covermode=count -coverprofile=cover.out

cover:
	@echo Opening coverage info in browser. If this failed run make test first
	$(GO) tool cover -html=cover.out
	$(GO) tool cover -func=cover.out
