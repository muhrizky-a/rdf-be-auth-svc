# Variables
PROTO_DIR=proto
GO_OUT_DIR=.
PROTOC=protoc
PROTOC_GEN_GO=protoc-gen-go
PROTOC_GEN_GO_GRPC=protoc-gen-go-grpc

# Find all .proto files
PROTO_FILES=$(wildcard $(PROTO_DIR)/*.proto)

# Generate Go code from .proto files
.PHONY: gen
gen:
	@mkdir -p $(GO_OUT_DIR)
	@$(PROTOC) --go_out=$(GO_OUT_DIR) --go-grpc_out=$(GO_OUT_DIR) --proto_path=$(PROTO_DIR) $(PROTO_FILES)
	@echo "Generated Go code from .proto files."

# Clean generated files
.PHONY: clean
clean:
	@rm -rf $(GO_OUT_DIR)
	@echo "Cleaned generated files."

# Help message
.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  gen    Generate Go code from .proto files."
	@echo "  clean  Remove generated files."
	@echo "  help   Show this help message."