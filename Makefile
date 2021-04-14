BIN_DIR := $(CURDIR)/.bin

$(BIN_DIR):
	@mkdir -p $(BIN_DIR)

.PHONY: lint
lint: golangci-lint
	golangci-lint run

golangci-lint: | $(BIN_DIR)
	@{ \
		set -e; \
		GOLANGCI_LINT_TMP_DIR=$$(mktemp -d); \
		cd $$GOLANGCI_LINT_TMP_DIR; \
		go mod init tmp; \
		GOBIN=$(BIN_DIR) go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.36.0; \
		rm -rf $$GOLANGCI_LINT_TMP_DIR; \
	}
