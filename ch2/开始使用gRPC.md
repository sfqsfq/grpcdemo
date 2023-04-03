- protocol buffer compiler(protoc) [protocolbuffers/protobuf: Protocol Buffers - Google's data interchange format (github.com)](https://github.com/protocolbuffers/protobuf#protocol-compiler-installation)
- protoc-gen-go: 生产 Go 代码 `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`
- protoc-gen-go-grpc: 生成 Go Grpc 代码 `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`


生成代码
```
protoc -I proto/ proto/product_info.proto --go_out=productinfo/service/server/ecommerce --go_opt=paths=source_relative --go-grpc_out=productinfo/service/server/ecommerce --go-grpc_opt=paths=source_relative

protoc -I proto/ proto/product_info.proto --go_out=productinfo/service/server/ecommerce --go_opt=paths=source_relative
```

