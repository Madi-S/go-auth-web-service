# To build and run

```bash
docker-compose up --build
```

## Prerequisites

1. protoc for Windows, protobuf-compiler for Linux.
2. grpcio-tools, grpclib, protobuf (python packages).
3. https://github.com/nipunn1313/mypy-protobuf (`pip install mypy-protobuf`)

### Run

```bash
go run cmd/sso/main.go
```

### Generate proto files

Run makefile:

```bash
make
```
