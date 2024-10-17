安装 buf
brew install bufbuild/buf/buf

go install github.com/bufbuild/buf/cmd/buf@v1.42.0

buf --version


初始项目：
buf config init

buf dep update

buf generate



github.com/golang/protobuf
废弃  同时生成pb和gRPC相关代码

google.golang.org/protobuf 
推荐使用
只生成pb序列化相关的文件
生成gRPC相关代码需要使用grpc-go插件protoc-gen-go-grpc

curl --header 'Content-Type: application/json' --data '{"ID": {"Value":22222}}' http://localhost:8080/rpc.User/Get

grpcurl -plaintext localhost:8080 list
grpcurl  -plaintext -d '{"ID": {"Value":22222}}' localhost:8080 rpc.User/Get