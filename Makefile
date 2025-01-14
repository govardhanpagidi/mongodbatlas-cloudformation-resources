tags="logging callback metrics scheduler"
cgo=0 
goos=linux
goarch=amd64

# Set next value at compile time for fallback log-level
ldXflags=github.com/mongodb/mongodbatlas-cloudformation-resources/util.defaultLogLevel=info
ldXflagsD=github.com/mongodb/mongodbatlas-cloudformation-resources/util.defaultLogLevel=debug

.PHONY: submit
submit:
	cd cfn-resources && ./cfn-submit-helper.sh $(filter-out $@,$(MAKECMDGOALS))

.PHONY: test
test:
	cd cfn-resources && ./cfn-testing-helper.sh $(filter-out $@,$(MAKECMDGOALS))

.PHONY: fmt
fmt: ## Format changed go and sh
	@scripts/fmt.sh

.PHONY: devtools
devtools:  ## Install dev tools
	@echo "==> Installing dev tools..."
	go install github.com/google/addlicense@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/google/go-licenses@latest
	go install mvdan.cc/sh/v3/cmd/shfmt@latest
	go install github.com/rhysd/actionlint/cmd/actionlint@latest
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin $(GOLANGCI_VERSION)

.PHONY: link-git-hooks
link-git-hooks: ## Install git hooks
	@echo "==> Installing all git hooks..."
	find .git/hooks -type l -exec rm {} \;
	find .githooks -type f -exec ln -sf ../../{} .git/hooks/ \;

.PHONY: lint
lint: ## Run linter
	@scripts/lint.sh
