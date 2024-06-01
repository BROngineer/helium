.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: go-imports
go-imports: goimports
	$(GOIMPORTS) -w .

.PHONY: vet
vet:
	go vet ./...

.PHONY: lint
lint: golangci-lint
	$(GOLANGCI_LINT) run

.PHONY: mod-tidy
mod-tidy:
	go mod tidy

.PHONY: test
test:
	go test ./... -coverprofile cover.out

LOCAL_BIN ?= $(shell pwd)/bin
$(LOCAL_BIN):
	mkdir -p $(LOCAL_BIN)

GOIMPORTS ?= $(LOCAL_BIN)/goimports
GOLANGCI_LINT ?= $(LOCAL_BIN)/golangci-lint

GOIMPORTS_VERSION ?= v0.16.1
GOLANGCI_LINT_VERSION ?= v1.56.2

.PHONY: golangci-lint
golangci-lint: $(GOLANGCI_LINT)
$(GOLANGCI_LINT): $(LOCAL_BIN)
	@test -s $(GOLANGCI_LINT) && $(GOLANGCI_LINT) --version | grep -q $(GOLANGCI_LINT_VERSION) || \
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

.PHONY: goimports
goimports: $(GOIMPORTS)
$(GOIMPORTS): $(LOCAL_BIN)
	@test -s $(GOIMPORTS) || GOBIN=$(LOCAL_BIN) go install golang.org/x/tools/cmd/goimports@$(GOIMPORTS_VERSION)
