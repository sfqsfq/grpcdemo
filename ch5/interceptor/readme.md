```
protoc -I proto/ proto/order_management.proto --go_out=ecommerce --go_opt=paths=source_relative --go-grpc_out=ecommerce --go-grpc_opt=paths=source_relative
```