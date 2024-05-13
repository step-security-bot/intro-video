.PHONY: test
test:
	@go test -v ./...

.PHONY: watch
watch:
	@air
