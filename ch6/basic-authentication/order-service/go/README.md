```
protoc --go_out=order-service-gen --go_opt=paths=source_relative --go-grpc_out=order-service-gen --go-grpc_opt=paths=source_relative -I  proto  order_management.proto
```