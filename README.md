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

### Generate proto files

Run makefile:

```bash
make
```

Update MANIFEST.in with following schema for each proto package:

```
recursive-include <package/folder name> *.proto
```

Install it as a symbolic link to package directory (just once is enough, because it is a local package):

```bash
cd protos/gen/python
pip install -e .
```
