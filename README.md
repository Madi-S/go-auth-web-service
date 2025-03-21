# To build and run

```bash
do something first lol
```

## Prerequisites

1. protoc for Windows, protobuf-compiler for Linux.
2. grpcio-tools, grpclib, protobuf (python packages).

### Run

```bash
go run cmd/sso/main.go
```

### Generate protos for go

```bash
protoc -I protos/proto protos/proto/<your proto file> --go_out=protos/gen/go --go_opt=paths=source_relative --go-grpc_out=protos/gen/go --go-grpc_opt=paths=source_relative
```

### Generate protos and a package for python

Generate proto files:

```bash
python -m grpc_tools.protoc -I protos/proto protos/proto/<your proto file> --python_out=protos/gen/python --grpc_python_out=protos/gen/python
```

Update MANIFEST.in with following schema for each proto package:

```
recursive-include <package/folder name> *.proto
```

Install it as a symbolic link to package directory:

```bash
cd protos/gen/python
pip install -e .
```

Or just use make and update MANIFEST.in by yourself

```bash
make
```
