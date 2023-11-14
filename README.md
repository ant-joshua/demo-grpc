

```bash
protoc \
--go_out=invoicer \
--go_opt=paths=source_relative \
--go-grpc_out=invoicer \
--go-grpc_opt=paths=source_relative \
invoicer.proto
```

This command is split indirect and direct dependencies into two separate block
```bash
go mod tidy
```