PROTO_SRC_DIR=protos
PYTHON_OUT_DIR=python/protos/gen
GO_OUT_DIR=go/protos/gen

# Find all .proto files
PROTO_FILES=$(wildcard $(PROTO_SRC_DIR)/**/*.proto)

.PHONY: all python go clean

all: python go

python:
	@echo "Generating Python Protobuf files..."
	python -m grpc_tools.protoc -I $(PROTO_SRC_DIR) $(PROTO_FILES) \
		--python_out=$(PYTHON_OUT_DIR) --grpc_python_out=$(PYTHON_OUT_DIR)

go:
	@echo "Generating Go Protobuf files..."
	protoc -I $(PROTO_SRC_DIR) $(PROTO_FILES) \
		--go_out=$(GO_OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(GO_OUT_DIR) --go-grpc_opt=paths=source_relative

clean:
	@echo "Cleaning generated files..."
	@if exist $(PYTHON_OUT_DIR) del /s /q $(PYTHON_OUT_DIR)\*.py
	@if exist $(GO_OUT_DIR) del /s /q $(GO_OUT_DIR)\*.go
