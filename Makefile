SHELL := /bin/bash

.PHONY: test
test:
	@go test -v ./...

.PHONY: watch
watch:
	@$(HOME)/go/bin/air & $(MAKE) tailwind-watch

.PHONY: templ
templ:
	@$(HOME)/go/bin/templ generate ./internal/template/*.templ

.PHONY: tailwind-watch
tailwind-watch:
	@bunx tailwindcss -i ./style.css -o ./public/style.css --watch

.PHONY: tailwind-build
tailwind-build:
	@bunx tailwindcss -i ./style.css -o ./public/style.css
