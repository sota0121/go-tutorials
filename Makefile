# Go Commands
GO = go
GO_RUN = $(GO) run
GO_BUILD = $(GO) build
GO_TEST = $(GO) test
GO_CLEAN = $(GO) clean
GO_GET = $(GO) get
GO_MOD = $(GO) mod
GO_FMT = $(GO) fmt
GO_VET = $(GO) vet
GO_TOOL = $(GO) tool

# Variables
MAIN_GO = cmd/main.go
BIN_DIR = bin
BIN_NAME = $(BIN_DIR)/tutorial
OUT_DIR = out
COVER_DIR = $(OUT_DIR)/cover
COVER_FILE = $(COVER_DIR)/cover.out
COVER_HTML = $(COVER_DIR)/cover.html

# Targets
.PHONY: all
all: clean test build

.PHONY: build
build:
	$(GO_BUILD) -o $(BIN_NAME) $(MAIN_GO)

.PHONY: test
test:
	@mkdir -p $(COVER_DIR)
	$(GO_TEST) -v -cover -coverprofile=$(COVER_FILE) ./...
	$(GO_TOOL) cover -html=$(COVER_FILE) -o $(COVER_HTML)

vet:
	$(GO_VET) ./...
fmt:
	$(GO_FMT) ./...

.PHONY: run
run:
	$(GO_RUN) $(MAIN_GO)

.PHONY: clean
clean:
	$(GO_CLEAN)
	@rm -rf $(BIN_DIR) $(OUT_DIR)
