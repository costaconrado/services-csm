run:
	go mod tidy && go mod download && \
	GIN_MODE=debug CGO_ENABLED=0 go run -tags migrate ./cmd/app
.PHONY: run

build-client:
	@echo "> Building the client..."
	yarn --cwd "./pkg/frontend_react" build
.PHONY: build-client