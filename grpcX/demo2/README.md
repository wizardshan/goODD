��װ buf
brew install bufbuild/buf/buf

go install github.com/bufbuild/buf/cmd/buf@v1.42.0

buf --version


��ʼ��Ŀ��
buf config init

buf dep update

buf generate



github.com/golang/protobuf
����  ͬʱ����pb��gRPC��ش���

google.golang.org/protobuf 
�Ƽ�ʹ��
ֻ����pb���л���ص��ļ�
����gRPC��ش�����Ҫʹ��grpc-go���protoc-gen-go-grpc

curl --header 'Content-Type: application/json' --data '{"ID": {"Value":22222}}' http://localhost:8080/rpc.User/Get

grpcurl -plaintext localhost:8080 list
grpcurl  -plaintext -d '{"ID": {"Value":22222}}' localhost:8080 rpc.User/Get